# Readme

Trying out json series mutation compressing with diffing.

Conclusion:

- Compressions aint good
- Unless its a wide dimension with mostly static fields, it wont give significant reduction than full json (changes metadata almost took same space like data)
- Also compression with zip almost gave same size (circa <3kb diff) with this method
