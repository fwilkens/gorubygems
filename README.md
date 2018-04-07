[![Build Status](https://travis-ci.org/fwilkens/gorubygems.svg?branch=master)](https://travis-ci.org/fwilkens/gorubygems)

## WIP Go wrapper for the [RubyGems API](http://guides.rubygems.org/rubygems-org-api/).

### Examples

Get the 50 most recently updated gems (50 is a non-param value decided by Rubygems).
```go
rubygemsapi.JustUpdated() // => []Gem
```

Search for a gem by keyword. Rubygems returns 30 records per page.
This wrapper doesn't have support for paging yet.
```go
rubygemsapi.Search("rails") // => []Gem
```

Get the 50 most recently created gems. (50 is a non-param value decided by Rubygems)
```go
rubygemsapi.Latest() // => []Gem
```

Get info for a particular gem
```go
rubygemsapi.GemInfo("rails") // => Gem
```

Gem structs have the following fields:
* `Name`
* `Version`
* `SourceCodeUri`
* `HomepageUri`
* `Platform`