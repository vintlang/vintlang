# Switch Statements in vint

Switch statements in vint allow you to execute different code blocks based on the value of a given expression. This page covers the basics of switch statements and their usage.

## Basic Syntax

You initialize a switch statement with the switch keyword, the expression inside parentheses (), and all cases enclosed within curly braces {}.

A case statement has the keyword ikiwa followed by a value to check. Multiple values can be in a single case separated by commas ,. The consequence to execute if a condition is fulfilled must be inside curly braces {}. Here's an example:

```s
let a = 2

switch (a){
	ikiwa 3 {
		print("a ni tatu")
	}
	ikiwa 2 {
		print ("a ni mbili")
	}
}
```

## Multiple Values in a Case

Multiple possibilities can be assigned to a single case (ikiwa) statement:

```s
switch (a) {
	ikiwa 1,2,3 {
		print("a ni kati ya 1, 2 au 3")
	}
	ikiwa 4 {
		print("a ni 4")
	}
}
```

## Default Case (kawaida)

The default statement will be executed when no condition is satisfied. The default statement is represented by kawaida:

```s
let z = 20

switch(z) {
	ikiwa 10 {
		print("kumi")
	}
	ikiwa 30 {
		print("thelathini")
	}
	kawaida {
		print("ishirini")
	}
}
```

By understanding switch statements in vint, you can create more efficient and organized code that can handle multiple conditions easily.