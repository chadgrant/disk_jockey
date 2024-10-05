package main

import (
	"strings"
	"testing"
)

var tests = []struct {
	songs []string
	tests []struct {
		start  string
		expect []string
	}
}{
	{
		songs: []string{
			"Down By the River", "River of Dreams", "Take me to the River", "Dreams", "Blues Hand Me Down", "Forever Young", "American Dreams", "All My Love", "Cantaloop", "Take it All", "Love is Forever", "Young American", "Dreamship", "Every Breath You Take",
		},
		tests: []struct {
			start  string
			expect []string
		}{
			{
				"Every Breath You Take",
				[]string{
					"Every Breath You Take", "Take it All", "All My Love", "Love is Forever", "Forever Young", "Young American", "American Dreams", "Dreams",
				},
			},
			{
				"Dreams",
				[]string{"Dreams"},
			},
			{
				"Blues Hand Me Down",
				[]string{"Blues Hand Me Down", "Down By the River", "River of Dreams", "Dreams"},
			},
			{
				"Cantaloop",
				[]string{"Cantaloop"},
			},
		},
	}, {
		songs: []string{"Bye Bye Love", "Nothing at All", "Money for Nothing", "Love Me Do", "Do You Feel Like We Do", "Bye Bye Bye", "Do You Believe in Magic", "Bye Bye Baby", "Baby Ride Easy", "Easy Money", "All Right Now"},
		tests: []struct {
			start  string
			expect []string
		}{
			{
				"Bye Bye Bye",
				[]string{"Bye Bye Bye", "Bye Bye Baby", "Baby Ride Easy", "Easy Money", "Money for Nothing", "Nothing at All", "All Right Now"},
			},
			{
				"Bye Bye Love",
				[]string{"Bye Bye Love", "Love Me Do", "Do You Feel Like We Do", "Do You Believe in Magic"},
			},
		},
	}, {
		songs: []string{"Love Me Do", "Do You Believe In Magic", "Magic You Do", "Magic Man", "Man In The Mirror"},
		tests: []struct {
			start  string
			expect []string
		}{
			{
				"Love Me Do",
				[]string{"Love Me Do", "Do You Believe in Magic", "Magic Man", "Man In The Mirror"},
			},
		},
	},
}

func TestPlaylists(t *testing.T) {
	for _, test := range tests {

		for _, tt := range test.tests {
			t.Run(tt.start, func(t *testing.T) {
				testPlaylist(t, tt.start, test.songs, tt.expect)
			})
		}
	}
}

func testPlaylist(t *testing.T, start string, songs []string, expect []string) {
	playlist := createPlaylist(start, songs)
	if len(playlist) != len(expect) {
		t.Errorf("Expected %d songs, got %d [%v]\n", len(expect), len(playlist), playlist)

		return
	}
	for i := 0; i < len(playlist); i++ {
		if !strings.EqualFold(playlist[i], expect[i]) {
			t.Errorf("Expected %s in position %d, got %s\n", expect[i], i, playlist[i])
			return
		}
	}
}
