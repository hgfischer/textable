* Get string length using proper unicode functions

* Load data from CSV
	* everything is text - try to discover what kind of information is inside strings

* Load data from DB
	* everything is typed

* Load data from HTML
	* everything is string, same as CSV

* Set ellipsis char …

* Set empty char (nil input)
	* ?/~/- 

* Set max width of each column
	* -1 = no max

* Store alignment for each column

* Sort entire table by a specific column, or "none" == nil
	* Table.SortBy(0) -> columns[0].Sort -> -> columns[x].Reorder(column[0].GetOrder))
	* Table.SortBy("Header")
	* Table.SortBy(nil) 
	* Table.ReverseSortBy(0)

* Test with UTF-8 strings

* Use fancy UTF-8 chars

* Return part of the table
	* Table.Limit(offset, limit)

* Different table styles
	* Style as a type
	* Table.SetStyle(style.MSWord)
	* Table.SetStyle(style.PlainColumns)