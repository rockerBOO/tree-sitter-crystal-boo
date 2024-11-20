package tree_sitter_crystal_boo_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_crystal_boo "github.com/rockerboo/tree-sitter-crystal-boo/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_crystal_boo.Language())
	if language == nil {
		t.Errorf("Error loading CrystalBoo grammar")
	}
}
