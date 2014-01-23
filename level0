#!/usr/bin/env ruby

# Our test cases will always use the same dictionary file (with SHA1
# 6b898d7c48630be05b72b3ae07c5be6617f90d8e). Running `test/harness`
# will automatically download this dictionary for you if you don't
# have it already.

class FastSearch
  def initialize(entries)
    @entries = entries.sort
  end

  def include?(word)
    start_point = 0
    end_point = entries.size
    mid_point = start_point + ((end_point - start_point)/2).truncate
    until start_point == mid_point || end_point == mid_point
      #puts start_point, mid_point, end_point, word

      case word <=> entries[mid_point]
      when -1
        end_point = mid_point
      when 0
        return true
      when 1
        start_point = mid_point
      end
    
      mid_point = start_point + ((end_point - start_point)/2).truncate

      return true if [entries[start_point], entries[mid_point], entries[end_point]].include?(word)
    end 
    return false
  end

  private
  attr_reader :entries
end

path = ARGV.length > 0 ? ARGV[0] : '/usr/share/dict/words'
fast_search = FastSearch.new(File.read(path).split("\n"))



contents = $stdin.read
output = contents.gsub(/[^ \n]+/) do |word|
  if fast_search.include?(word.downcase)
    word
  else
    "<#{word}>"
  end
end
print output
