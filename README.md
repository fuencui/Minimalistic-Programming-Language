# Minimalistic-Programming-Language
Minimalistic Programming Language; Racket style built by golang

Welcome to minimalistic racket phase!
> (= (+ 1 2) 3)
[FINAL RESULT]: #t
> (= (+ 1 2) (- 4 3))
[FINAL RESULT]: #f
> (>= (+ 1 2) (- 3 (* 1 2)))
[FINAL RESULT]: #t
> (and (= 1 2) (> (/ 1 0) 3))
[FINAL RESULT]: #f
> (and (not (= 1 2)) (or (> 1 2) (= (+ 3 4) 7)))
[FINAL RESULT]: #t
> (if (and (> 1 2) (= 3 4)) (/ 1 0) (or false (if true (= 1 2) (/ 4 5))))
[FINAL RESULT]: #f
> (define x (+ 1 2))
> x
[FINAL RESULT]: 3.000000
> (define y (* x x))
> y
[FINAL RESULT]: 9.000000
> z
[FINAL RESULT]: missing a variable error
> (define z (+ z 1))
> z
[FINAL RESULT]: cannot reference an identifier before its definition
> (define (fib n) (if (< n 2) n (+ (fib (- n 1)) (fib (- n 2)))))
> (fib 0)
[FINAL RESULT]: 0.000000
> (fib 1)
[FINAL RESULT]: 1.000000
> (fib 2)
[FINAL RESULT]: 1.000000
> (fib x)
[FINAL RESULT]: 2.000000
> (fib (+ x 1))
[FINAL RESULT]: 3.000000
> (fib (+ (* 2 x) y))
[FINAL RESULT]: 610.000000
> (define (addx a) (+ a x))
> (addx 12)
[FINAL RESULT]: 15.000000
> (define x 5)
> (addx 12)
[FINAL RESULT]: 17.000000
> (define (xplusy) (+ x y))
> (xplusy)
[FINAL RESULT]: 30.000000
> (define x 11)
> (define y 22)
> (xplusy)
[FINAL RESULT]: 33.000000
> (define (add1 x) (+ x 1))
> (define (add2 x) (add1 (add1 x)))
> (add1 12)
[FINAL RESULT]: 13.000000
> (add2 12)
[FINAL RESULT]: 14.000000

