package parse

import (
	"dataframe/vendor_lib/json2jsonschema/lex"
	"errors"
	"fmt"
)

type parser struct {
	Lexer    *lex.Lexer
	Item     lex.Item
	LastItem lex.Item
	MapFlat  map[string]string
	//MapType  map[string]string
}

func JsonStringToFlatMap(jsonString string) (map[string]string, error) {
	parse := &parser{
		Lexer:   lex.Lex("", jsonString),
		MapFlat: map[string]string{},
	}
	return parse.parseToFlatMap()
}

func (p *parser) parseToFlatMap() (_ map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			//node = nil
			err = errors.New(fmt.Sprint(r))
		}
	}()

	p.Item = p.Lexer.NextItem()

	switch p.Item.Typ {
	case lex.ItemLeftBrace:
		p.parseObject("")
		return p.MapFlat, nil
	//case lex.ItemLeftSqrBrace:
	//	return p.parseArray(), nil
	case lex.ItemError:
		panic(fmt.Sprintf("received error from lexer at pos %v: %v", p.Item.Pos, p.Item.Value))
	default:
		panic(fmt.Sprintf("error determining root json type. unexpected item %v", p.Item.Value))
	}
}

func (p *parser) parseObject(rootKey string) {
	p.LastItem = p.Item
	if len(rootKey) > 0 {
		rootKey += "."
	}

	var currentKey string
	for p.Item = p.Lexer.NextItem(); p.Item.Typ != lex.ItemRightBrace; p.Item = p.Lexer.NextItem() {

		switch p.Item.Typ {
		case lex.ItemString:

			// A string can indicate that the current lexem is either a key or a value.
			// It's a key if the previous lexem is a comma or an opening brace.
			// It's a value if the previous lexem is a colon.
			if p.LastItem.Typ == lex.ItemComma || p.LastItem.Typ == lex.ItemLeftBrace {
				currentKey = p.Item.Value
			} else {
				p.MapFlat[rootKey+currentKey] = p.Item.Value
				//p.MapType[rootKey+currentKey] = "string"
			}
		case lex.ItemLeftBrace:
			p.parseObject(rootKey + currentKey)
		case lex.ItemLeftSqrBrace:
			p.parseArray(rootKey + currentKey)
		case lex.ItemBool:
			p.MapFlat[rootKey+currentKey] = p.Item.Value
		case lex.ItemNil:
			p.MapFlat[rootKey+currentKey] = p.Item.Value
		case lex.ItemFloat:
			p.MapFlat[rootKey+currentKey] = p.Item.Value
		case lex.ItemInteger:
			p.MapFlat[rootKey+currentKey] = p.Item.Value
		case lex.ItemColon:
			break
		case lex.ItemComma:
			break
		case lex.ItemError:
			panic(fmt.Sprintf("received error from lexer. pos: %v, msg: %v", p.Item.Pos, p.Item.Value))
		default:
			panic(fmt.Sprintf("error parsing object. unexpected item %v", p.Item.Value))
		}
		p.LastItem = p.Item
	}

	if p.LastItem.Typ == lex.ItemComma {
		panic("error parsing object. a closing curly brace mustn't follow a comma")
	}

	return
}

func (p *parser) parseArray(rootKey string) {
	p.LastItem = p.Item

	for p.Item = p.Lexer.NextItem(); p.Item.Typ != lex.ItemRightSqrBrace; p.Item = p.Lexer.NextItem() {
		switch p.Item.Typ {
		case lex.ItemLeftBrace:
			p.parseObject(rootKey)
		case lex.ItemLeftSqrBrace:
			p.parseArray(rootKey)
		case lex.ItemNil:
			p.MapFlat[rootKey] += p.Item.Value + ","
		case lex.ItemBool:
			p.MapFlat[rootKey] += p.Item.Value + ","
		case lex.ItemString:
			p.MapFlat[rootKey] += p.Item.Value + ","
		case lex.ItemInteger:
			p.MapFlat[rootKey] += p.Item.Value + ","
		case lex.ItemFloat:
			p.MapFlat[rootKey] += p.Item.Value + ","
		case lex.ItemComma:
			break
		case lex.ItemError:
			panic(fmt.Sprintf("received error from lexer. pos: %v, msg: %v", p.Item.Pos, p.Item.Value))
		default:
			panic("error parsing array: unexpected item " + p.Item.Value)
		}
		p.LastItem = p.Item
	}

	if p.LastItem.Typ == lex.ItemComma {
		panic("error parsing array: a closing square brace mustn't follow a comma")
	}
}
