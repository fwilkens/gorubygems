## WIP Go wrapper for the [RubyGems API](http://guides.rubygems.org/rubygems-org-api/).

### Examples

Get the 50 most recently updated gems (50 is a non-param value decided by Rubygems).

Get the 50 most recently created gems. (50 is a non-param value decided by Rubygems)

Search for a gem by keyword. Rubygems returns 30 records per page.
This wrapper doesn't have support for paging yet.
```go
rubygemsapi.JustUpdated() // => []Gem
```

```go
rubygemsapi.Search() // => []Gem
```

```go
rubygemsapi.Latest() // => []Gem
```

Gem structs have the following fields:
* `Name`
* `Version`
* `SourceCodeUri`
* `HomepageUri`
* `Platform`