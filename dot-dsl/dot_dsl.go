package dotdsl

import (
	"errors"
	"maps"
	"regexp"
	"strconv"
	"strings"
)

// Properties holds the properties of a node or edge.
// The values can be int, bool or string.
type Properties map[string]any

// Graph stores the parts of a dot graph.
// All entities are stored as a Properties map (`nil` Properties when none set)
// attrs is the Properties for the entire Graph, vs a specific node or edge.
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}

// Tokens keep quoted values intact and retain newlines as statement separators.
var graphTokens = regexp.MustCompile(`//[^\n]*|#[^\n]*|"(?:\\[^\n]|[^"\\\n])*"|--|[+-]?[0-9]+\b|[[:alnum:]_]+|[^ \t\r\n]|\n`)
var nodeName = regexp.MustCompile(`^[[:alnum:]]+$`)

var (
	errGraph = errors.New("invalid graph")
	errEdge  = errors.New("invalid edge")
	errAttr  = errors.New("invalid attribute")
	errNode  = errors.New("node name must be alphanumeric")
)

type parser struct {
	tokens []string
}

func (p *parser) peek() string {
	if len(p.tokens) == 0 {
		return ""
	}
	return p.tokens[0]
}

func (p *parser) take() string {
	token := p.peek()
	if token != "" {
		p.tokens = p.tokens[1:]
	}
	return token
}

func (p *parser) accept(token string) bool {
	if p.peek() != token {
		return false
	}
	p.take()
	return true
}

func (p *parser) skipNewlines() {
	for p.accept("\n") {
	}
}

// Parse creates a Graph from a text blob.
func Parse(data string) (*Graph, error) {
	tokens := graphTokens.FindAllString(data, -1)
	p := parser{tokens: tokens[:0]}
	for _, token := range tokens {
		if !strings.HasPrefix(token, "//") && !strings.HasPrefix(token, "#") {
			p.tokens = append(p.tokens, token)
		}
	}

	p.skipNewlines()
	if !p.accept("graph") {
		return nil, errGraph
	}
	p.skipNewlines()
	if !p.accept("{") {
		return nil, errGraph
	}

	graph := &Graph{
		nodes: make(map[string]Properties),
		edges: make(map[string]Properties),
	}
	for {
		p.skipNewlines()
		switch p.peek() {
		case "}":
			p.take()
			p.skipNewlines()
			if p.peek() != "" {
				return nil, errGraph
			}
			return graph, nil
		case "":
			return nil, errGraph
		case ";":
			p.take()
		default:
			if err := p.statement(graph); err != nil {
				return nil, err
			}
		}
	}
}

func (p *parser) statement(graph *Graph) error {
	var names []string
	if p.peek() != "[" {
		name := p.take()
		if name != "graph" || p.peek() != "[" {
			if !nodeName.MatchString(name) {
				return errNode
			}
			names = append(names, name)
			for p.accept("--") {
				p.skipNewlines()
				name = p.take()
				if !nodeName.MatchString(name) {
					return errEdge
				}
				names = append(names, name)
			}
		}
	}

	props, err := p.properties()
	if err != nil {
		return err
	}
	switch p.peek() {
	case ";", "\n", "}":
	default:
		return errEdge
	}

	switch len(names) {
	case 0:
		graph.attrs = mergeProperties(graph.attrs, props)
	case 1:
		graph.nodes[names[0]] = mergeProperties(graph.nodes[names[0]], props)
	default:
		for i, name := range names {
			if _, exists := graph.nodes[name]; !exists {
				graph.nodes[name] = nil
			}
			if i > 0 {
				a, b := names[i-1], name
				if a > b {
					a, b = b, a
				}
				key := "{" + a + " " + b + "}"
				graph.edges[key] = mergeProperties(graph.edges[key], props)
			}
		}
	}
	return nil
}

func (p *parser) properties() (Properties, error) {
	var props Properties
	for p.accept("[") {
		for {
			p.skipNewlines()
			key := p.take()
			p.skipNewlines()
			if !nodeName.MatchString(key) || !p.accept("=") {
				return nil, errAttr
			}
			p.skipNewlines()
			value, err := propertyValue(p.take())
			if err != nil {
				return nil, err
			}
			if props == nil {
				props = make(Properties)
			}
			props[key] = value
			p.skipNewlines()
			if p.accept("]") {
				break
			}
			if !p.accept(",") && !p.accept(";") {
				return nil, errAttr
			}
		}
	}
	return props, nil
}

func propertyValue(token string) (any, error) {
	switch token {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	if strings.HasPrefix(token, `"`) {
		value, err := strconv.Unquote(token)
		if err != nil {
			return nil, errAttr
		}
		return value, nil
	}
	if value, err := strconv.Atoi(token); err == nil {
		return value, nil
	}
	if nodeName.MatchString(token) {
		return token, nil
	}
	return nil, errAttr
}

func mergeProperties(existing, update Properties) Properties {
	if existing == nil {
		return maps.Clone(update)
	}
	maps.Copy(existing, update)
	return existing
}
