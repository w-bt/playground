# Christmas Lights Pattern Problem

A string of Christmas lights needs to be programmed with a repeating pattern as shown below:

Implement the `next` method in the `ChristmasLights` class that returns an array of booleans to display the pattern.

The `next` method is called to display the next state in the pattern, and this repeats indefinitely. 
In the example video shown above, the first time the `next` method is called, it returns the following array:

```python
[true, true, false, false, false, false, false, false, true]
```

where `true` indicates that the light is on and `false` means that the light is off. 
There will be at least 4 lights on the string.
