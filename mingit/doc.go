// Package mingit is a utility library for programmatically generating
// minimally-functional local Git repositories.
//
// This is intended for the narrow use-case of using a Git repository to
// represent some data generated entirely by a program, such as a log of
// changes made to some data objects over the course of several steps. The
// goal is only to generate as easily as possible a bare-bones git repository
// directory with just enough content for normal git tools to be able to
// understand it.
//
// There are lots of things that mingit doesn't do, including but not limited
// to: reading from existing repositories, generating pack files and index
// files, and interacting with git's network protocols. If you need git
// functionality not available here then consider using
// github.com/go-git/go-git instead.
package mingit
