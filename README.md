# authentication

### General Thoughts
A large part of the challenge for me has been figuring out how the language works, how to get a Go environment running, 
particularly figuring out the imports (I would have preferred to use the github.com/golang-jwt/jwt package but I didn't realise I had imported it before I had solved it with jose)

### Language used and reason
I used Go, it seemed smart to learn the thing I'll hopefully be working with in the future.
It's been fun to use Go in a way that wasn't the go.dev playground.
### Example signature and decoded output
The code generates the following JWT:
```
eyJhbGciOiJIUzI1NiJ9.CnsKICAgICJuYW1lIjogT2JpLVdhbiBLZW5vYmksCiAgICAi
ZW1haWwiOiBiZW5Aa2Vub2JpLmNvbSwKICAgICJyb2xlcyI6IEplZGkgTWFzdGVyCn0K.pG5zLr5NYPT
Wa5kAbsW6VBDQsDHQ87m1l3MMt1mAzRA
```
Which should decode to:
```
{
"name": Obi-Wan Kenobi,
"email": ben@kenobi.com,
"roles": Jedi Master
}
``` 
