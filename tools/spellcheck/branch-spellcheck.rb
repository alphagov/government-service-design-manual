#!/usr/bin/env ruby
# encoding: UTF-8

# Usage: ./tools/spellcheck/branch-spellcheck.rb name-of-branch

require 'nokogiri'

if ARGV[0].nil?
  puts "No branch name specified. Exiting."
  exit 0
end

puts "-----------------------------"
puts "Checking changed files for spelling errors..."
puts "This is a very basic script and needs huge amounts of care. Do not rely on it to be perfectly correct."
puts "Spellchecking is hard."
puts "It won't fail the build."
puts "-----------------------------"

# Get a list of files that have been changed on this branch

files_changed = `git diff --name-only -- master...#{ARGV[0]}`.split("\n")

# If the file is in the service-manual/ directory and ends in *.md, process it

files_changed.delete_if do |filename|
  not (filename.start_with? 'service-manual/' and filename.end_with? '.md')
end

files_changed.map! do |filename|
  "_site/#{filename}".gsub('.md', '.html')
end

# Use the compiled Jekyll version because we want to test the thing that will
# actually be displayed to the user.

# We could do this by indicating to the spellchecker that the input document is
# SGML/HTML, but curly quotes will make you want to run around screaming for a
# while. Then you'll start reading about how spellcheckers work and you'll cry.

words_added = []

files_changed.each do |filename|

  f = File.open(filename)
  doc = Nokogiri::HTML(f)

  # Body text
  words_added.concat(doc.css('p').map(&:text).join("\n").split("\s"))

  # Headings
  words_added.concat(doc.css('h1').map(&:text).join("\n").split("\s"))
  words_added.concat(doc.css('h2').map(&:text).join("\n").split("\s"))
  words_added.concat(doc.css('h3').map(&:text).join("\n").split("\s"))

  # TODO: ordered and unordered lists?

  f.close
end

words_added.map! do |word|
  # Replace curly right single quotes with straight ones so that aspell can do its thing
  word.gsub("’", "'")
end

wordfile = File.open('tmp/spellcheck.txt', 'w')
wordfile.write(words_added.join("\n"))
wordfile.close

puts `cat tmp/spellcheck.txt | aspell --list --master=british --personal=./tools/spellcheck/gds-words.txt`

# This is so flaky that we can't let it fail the build.
# Right now it should be used as a reviewing tool in addition to your eyes.
exit 0
