package main

import (
	"slices"
	"sort"
	"strings"
)

type Node struct {
	Song     string
	Parent   *Node
	Children []*Node
}

func (n *Node) HasSong(s string) bool {
	if n.Song == s {
		return true
	}

	for parent := n.Parent; parent != nil; parent = parent.Parent {
		if parent.Song == s {
			return true
		}
	}

	return false
}

func createPlaylist(start string, songs []string) []string {
	root := tree(start, songs)
	leaves := leaves(root)

	var out [][]string
	for i := 0; i < len(leaves); i++ {
		playlist := buildPlaylistFromLeaf(leaves[i])
		out = append(out, playlist)
	}

	sort.Slice(out, func(i, j int) bool {
		return len(out[i]) > len(out[j])
	})

	return out[0]
}

func buildPlaylistFromLeaf(n *Node) []string {
	out := make([]string, 0)

	t := n
	for {
		if t.Parent == nil {
			out = append(out, t.Song)
			break
		}

		out = append(out, t.Song)
		t = t.Parent
	}

	slices.Reverse(out)

	return out
}

func tree(start string, songs []string) *Node {
	n := &Node{Song: start}
	n.Children = findEndingIn(n, songs)
	return n
}

func leaves(n *Node) []*Node {
	if len(n.Children) == 0 {
		return []*Node{n}
	}

	out := make([]*Node, 0)

	for _, c := range n.Children {
		if len(c.Children) == 0 {
			out = append(out, c)
		} else {
			out = append(out, leaves(c)...)
		}
	}

	return out
}

func findEndingIn(parent *Node, songs []string) []*Node {
	children := make([]*Node, 0)

	parts := split(parent.Song)
	last := parts[len(parts)-1]

	for i := 0; i < len(songs); i++ {
		if (songs[i] == last && parent.Song == songs[i]) || parent.HasSong(songs[i]) {
			continue
		}
		if split(songs[i])[0] == last {
			children = append(children, &Node{Parent: parent, Song: songs[i]})
		}
	}

	for _, c := range children {
		c.Children = findEndingIn(c, songs)
	}

	return children
}

func split(s string) []string {
	return strings.Split(s, " ")
}
