Yet another tagged filesystem
=============================

Usage examples
--------------

### Documents tagging

Assume you working with company named foo and create some documents

	$ touch foo/agreement.doc
	$ touch foo/invoice.pdf
	$ touch foo/proposal.xls

And also you work with other company
	
	$ touch bar/agreement.doc
	$ touch bar/invoice.pdf
	$ touch bar/proposal.xls

Now you can list files either by company name
	
	$ ls foo/
	agreement.doc
	invoice.pdf
	proposal.xls

either by document type (TODO: would it work like that(?)

	$ ls agreement/
	foo.doc
	bar.doc

### Todo tagging

And now you want to tag documents for todo

	$ mv foo/agreement.doc foo/todo.agreement.doc
	$ mv bar/invoice.pdf bar/todo.invoice.pdf

Now you can list all the todo items

	$ ls todo/
	foo.agreement.doc
	bar.invoice.pdf

### Photos tagging

TODO

### Torrent tagging

TODO
