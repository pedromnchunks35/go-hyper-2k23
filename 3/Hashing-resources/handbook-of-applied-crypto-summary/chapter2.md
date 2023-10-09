# Mathematical Background
## Notations
- **Z** denotes the set of integers
- **Q** denotes the set of rational numbers
- **R** denotes the set of real numbers
- π = 3,14159
- **e** is the base of the natural logarithm; **e**~=2,17828
- [**a**,**b**] denotes the intefers **x** satisfying **a**<=**x**<=**b**
- [**x**] is the largest integer less than or equal to **x**. Ex: [5.2] = 5 and [-5.2] = -6
- [**x**] is the smallest integer greater than or equal to **x**. Ex: [5.2] = 6 and [-5.2] = -5
- If **A** is a finite set, then |**A**| denotes the number of elements in **A**, called the cardinality of **A**
- **a** E **A** means that element **a** is a member of the set **A**
- **A** C/ **B** means that **A** is a subset of **B**
- **A** C **B** means that **A** is a subset of **B**, that is **A** C/ **B** and **A** != **B**
- Intersection of sets is **A** ∩ **B**
- Union of sets is **A** ∪ **B**
- The difference of sets is **A** - **B**
- The cartesian product of sets **A** and **B** is the set **A** X **B** = {(**a**,**b**)|**a** E **A** and **b** E **B**}. Example: {**a_1**,**a_2**} X {**b_1**,**b_2**,**b_3**} = {(**a_1**,**b_1**),(**a_1**,**b_2**),(**a_1**,**b_3**),(**a_2**,**b_1**),(**a_2**,**b_2**),(**a_2**,**b_3**)}
- **f**: **A** -> **B**, means the f(**a**)=**b**, where **a** E **A** and **b** E **B**. **A** is the domain(**a** is the object) and **B**(**b** is the image) the codomain
- **f**: **A** -> **B** is 1-1 or injective if each element of **B** is the image of at most one element in **A**
- **f**: **A** -> **B** is onto or surjective if each **b** E **B** is the image of at least one **a** E **A**
- **f**: **A** -> **B** is a bijection if it is both 1-1 and onto. if **f** is a bijection between finite sets **A** and **B**, then |**A**|=|**B**|(ex: A{1,2,3} B(10,9,11)). If **f** is a bijection between set **A** and itself, then **f** is called a permutation on **A**(ex: f(1)=2 f(2)=3 f(3)=1, the domain is **A** and also the codomain is **A**)
- **ln**(**x**) is the natural logarithm of **x**; that is the logarithm of **x** to the base **e**
- **lg**(**x**) is the logarithm of **x** to the base 2
- **exp**(**x**) is the exponential function **e**^(**x**)
- The expression bellow denotes the sum **a_1**+**a_2**+...+**a_n**
  ```
  / n
  |     a_i
  \ i=1
  ```
- The following expression denotes the produc **a_1** x **a_2** x ... x **a_n**
  ```
  __ n
  ||    a_i
    i=1
  ```
- for a positive integer **n**, the factorial function is **n**! = **n**(**n**-1)(**n-2**)...1. 0! = 1
## Probability Theory
### Definition 
- An **experiment** is a procedure that yields one of a given set of outcomes. 
- **simple events** are individual possible outcomes
- **sample space** is the set of all possible outcomes
- **discrete sample spaces** are the most common samples in this chaper and it is **sample spaces** with only finitely many possible outcomes
- Simple events of a sample space **S** will be labeled **S_1**,**S_2**,...,**S_n**
### Definition 
- A **probability distribution** **P** on **S** is a sequence of numbers **p_1**,**p_2**,...,**p_n** that are all non-negative and sum to 1
- The number **p_i** is interpreted as the probability of **s_i** being the outcome of the experiment
### Definition 
- An event **E** is a subset of the sample space **S**
- The probability that event **E** occurs, denoted **P**(**E**),is the sum of the probabilities **p_i** of all simple events **s_i** which belong to **E**
- If **s_i** E **S**,**P**({**s_i**}) is simply denoted by **P**(**s_i**)
### Definition
- If **E** is an event, the **complementary event** is the set of simple events not belong to **E**, denoted **E**^-
### Fact
- Let **E** ⊆ **S**
- 0 <= **P**(**E**) <= 1, **P**(**S**) = 1 and **P**(**θ**) = 0 (**θ** is the empty set)
- **P**(**E**^-) = 1 - **P**(**E**)
- If the outcomes in **S** are equally likely, then **P**(**E**) = |**E**| / |**S**|
### Definition
- Two events **E_1** and **E_2** are called **mutually exclusive** if **P**(**E_1**∩**E_2**)=0
- This means that by one happening the other is ommited
### Fact
- Let **E_1** and **E_2** be two events
- If **E_1** ⊆ **E_2** then **P**(**E_1**) <= **P**(**E_2**)
- **P**(**E_1**∪**E_2**)+**P**(**E_1**∩**E_2**)=**P**(**E_1**)+**P**(**E_2**)
- Case **E_1** and **E_2** are mutually exclusive, then **P**(**E_1**∪**E_2**)=**P**(**E_1**)+**P**(**E_2**)
## Conditional Probability
### Definition
- **E_1** and **E_2** two events with **P**(**E_2**) > 0
- The **conditional probability** of **E_1** given **E_2**,denoted **P**(**E_1**|**E_2**) is **P**(**E_1**|**E_2**)=**P**(**E_1**∩**E_2**)/**P**(**E_2**)
- **P**(**E_1**|**E_2**) means the probability of event **E_1** occurring, given that **E_2** has ocurred
### Definition
- Events **E_1** and **E_2** are **independent** if **P**(**E_1**∩**E_2**)=**P**(**E_1**)x**P**(**E_2**)
- If **E_1** and **E_2** are independent, then **P**(**E_1**|**E_2**)=**P**(**E_1**) and **P**(**E_2**|**E_1**)=**P**(**E_2**), this means that the occurence of one does not influence the ocurrence of the another
### Fact
- If **E_1** and **E_2** are events with **P**(**E_2**)>0, then **P**(**E_1**|**E_2**)=(**P**(**E_1**)x**P**(**E_2**|**E_1**))/**P**(**E_2**)
## Random variables
- Let **S** be a sample space with probability distribution **P**
### Definition
- A random variable **X** is a function from the sample space **S** to the set of real numbers
- to each simple event **s_i** E **S**
- **X** assigns a real number **X**(**s_i**)
- Since **S** is assumed to be finite, **X** can only take on a finite number of values
### Definition
- **X** is a random variable on **S**
- The **expected value** is:
  ```
  /
  | X(s_i)xP(s_i)
  \ 
    s_i E S
  ```
### Fact
- **X** is a random variable on **S**
- Then:
  ```
  /
  | (x) x P(X=(x))
  \
    x E R
  ```
### Fact
- If **X_1**,**X_2**,...,**X_m** are random variables on **S**
- And **a_1**,**a_2**,...,**a_m** are real numbers, then:
  ```
        m                m
    ( /           )    /
  E( |  a_i x X_i ) = | a_i x E(X_i)
    ( \           )    \
        i=1              i=1 
  ```
### Definition
- The variance of a random variable **X** of mean **u** is a non-negative number **VAR**(**X**)=**E**((**X**-**u**)^2)
- The standard deviation of **X** is the non-negative square root of **Var**(**X**)
- If a random variable has small variance then large deviations from the mean are unlikely to be observed
### Fact
- **X** is a random variable with mean **u** = **E**(**X**) and variance **o**^2 = **VAR**(**X**)
- Then for any **t** > 0:
  ```
  P(|X-u| >= t) <= ((o^2)/(t^2))
  ```
## Binomial Distribution
### Definition
- **n** and **k** are non-negative integers
- The binomial coefficient $\binom{n}{k}$
- is the number of diferent ways of choosing **k** distinct objects from a set of **n** distinct objects
- the order is not important
### Fact (properties of binomial coefficients)
- **n** and **k** are non-negative integers
1. $\binom{n}{k}$ = $\frac{n!}{k!*(n-k)!}$  
2. $\binom{n}{k}$ = $\frac{n}{n-k}$
3. $\binom{n+1}{k+1}$ = $\binom{n}{k}$ + $\binom{n}{k+1}$
```
(
  n+1 = (         (
  k+1     n   +    n
)         k        k+1  
        )          ) 
```
### Fact (binomial theorem)
- For any real numbers **a**,**b** and non-negative integer **n**
- $(a+b)^n$=$\sum_{k=0}^{n}$$\binom{n}{k}$$*$$a^{k}$$*$$b^{n-k}$
### Bernoulli trial
- Experiment with exacly two possible outcomes
- These two possible outcomes are called success and failure
### Fact
- Suppose that the probability of success on a particular Bernoulli trial is **p**
- Then the probability of exacly **k** successes in a sequence of **n** such independent trials is: $$\binom{n}{k}$$ $$p^k$$ $$(1-p)^{n-k}$$, for each $$0 \le k \le n$$
### Definition
- The probability distribution is called **binomial distribution**
### Fact
- The expected number of successes in a sequence of **n** independent Bernoulli trials, with probability **p** of success in each trial, is **np**
- The variance of the number of successes is **np**(1-**p**)
### Fact (law of large numbers)
- **X** is a random variable denoting the fraction of sucesses in **n** independent Bernoulli trials
- **p** is the probability of success in each trial
- Then for any **e** > 0, **P**(|**X**-**p**|$>$**e**) -> 0, as **n** -> $\infty$
- In other words, as **n** gets larger, the proportion of successes should be close to **p**
## Birthday problems
1.
- For positive integers **m**,**n**
- With **m**>=**n**
- The number $m^{(n)}$ = **m**(**m**-1)(**m**-2)...(**m**-**n**+1)
2.
- **m**,**n** are non-negative integers with **m** >= **n**
- The stirling number of the second kind, denoted like $\left\{\frac{m}{n}\right\}$ is: $\left\{\frac{m}{n}\right\}$ = $\frac{1}{n!}$  $*$ $\sum_{k=0}^{n}$ $*$ $(-1)^{n-k}$ $*$ $\binom{n}{k}$ $*$ $k^{m}$
- The exception is that $\left\{\frac{0}{0}\right\}$ = 1
- $\left\{\frac{m}{n}\right\}$ counts the number of ways of partioning a set of **m** objects into **n** non-empty subsets
### Fact (classic occupancy problem)
- An urn has **m** balls numbered 1 to **m**
- **n** balls are drawn from the urn one at a time, with replacement, and they numbers are listed
1.
- The probability of at least one coincidence (a ball drawn at least twice) is: $P_{1}$(m,n,t) = $\binom{n}{t}$ $*$ $\frac{m^{(t)}}{n^{(n)}}$, $1 \le t \le n$
### Fact (birthday problem)
- An urn has **m** balls numbered 1 to **m**
- **n** balls are drawn from the urn one at a time. With replacement and the numbers are listed
1.
- The probability of at least one coincidence (a ball get drawn twice) is: $P_2(m,n)$=1-$P_1(m,n,n)$=1-$\frac{m_{(n)}}{m_n}$,$1 \le n \le m$
- If n=O($\sqrt{m}$) and m -> $\infty$
- Then $P_2$(m,n) -> 1 - exp($\frac{-n(n-1)}{2*m}$+O($\frac{1}{\sqrt{m}}$)) $\approx$ 1 - exp($\frac{-n^{2}}{2*m}$)
2.
- As m->$\infty$, the expected number of draws before a coincidence is $\sqrt{\frac{\pi*m}{2}}$
#### Explanation (why probability distribution is refered as birthday surprise or birthday paradox)
- The probability of 2 people in a room of 23 people having the same birthday is $P_{2}$(365,23) $\approx$ 0.507
- This number is "large"
- $P_2$(365,n) increases rapidly as **n** increases
### Facts
- There are 2 urns
- One containing **m** white balls numbered 1 to **m**
- The other contains **m** red balls numbered 1 to **m**
- First, $n_1$ balls are selected from the first urn and their numbers listed
- Then $n_2$ balls are selected from the second urn and their numbers listed
- The number of **coincidences** between the two lists are counted
#### Model A
- In case the balls are drawn one at a time
- With replacement
- The probability of at least one coincidence is $P_3$(m,$n_1$,$n_2$) = 1 - $\frac{1}{m^{n_1+n_2}}$ * $\sum_{t_1,t_2}$$m^{t_1+t_2}$ $\left\{\frac{n_1}{t_2}\right\}$ $\left\{\frac{n_2}{t_2}\right\}$. The summation is over all $0 \le t_1 \le n_1$,$0 \le t_2 \le n_2$
- Case the n=$n_1$=$n_2$,n=O($\sqrt{m}$) and m->$\infty$
- then $P_3$(m,$n_1$,$n_2$) -> 1 - exp(-$\frac{n^{2}}{m}$ $*$ [1+O($\frac{1}{\sqrt{m}})$]) $\approx$ 1 - exp(-$\frac{n^{2}}{m}$)
### Model B
- If the balls from both urns are drawn without replacement, then the probability of at least one coincidence is is $P_4$(m,$n_1$,$n_2$)=1-$\frac{m^{(n_1+n_2)}}{m^{(n_1)}*m^{(n_2)}}$.
- If $n_1$=O($\sqrt{m}$),$n_2$=O($\sqrt{m}$),and m -> $\infty$
- Then $P_4$(m,$n_1$,$n_2$)->1-exp(-$\frac{n_1*n_2}{m}$*[1+$\frac{n_1+n_2-1}{2*m}$+O($\frac{1}{m}$)])
### Model C
- If the $n_1$ white balls are drawn one at a time, with replacement, and the $n_2$ red balls are drawn without replacement
- Then the probability of at least one coincidence is: $P_5$(m,$n_1$,$n_2$)=1-$(1-\frac{n_2}{m})^{n_1}$
- If $n_1$ = O($\sqrt{m}$),$n_2$=O($\sqrt{m}$) and m->$\infty$ then $P_5$(m,$n_1$,$n_2$) -> 1 - exp(-$\frac{n_1*n_2}{m}$*[1+O($\frac{1}{\sqrt{m}})$]) $\approx$ 1-exp(-$\frac{n_1*n_2}{m}$)
## Random mappings
- Let $F_n$ denote the collection of all functions (mappings) from a finite domain of size **n** to a finite codomain of size **n**
- Models where random elements of $F_n$ are considered are called **random mappings models**
- such models arise frequently in cryptography and algorithmic number theory
- |$F_n$|=$n^{n}$
- The probability that a particular function from $F_n$ is chosen is $\frac{1}{n^{n}}$
### Another definition
- Let **f** be a function in $F_n$ with domain and codomain equal to {1,2,...,**n**}
- The functional graph of **f** is a directed graph whose points (or vertices) are the elements {1,2,...,**n**} and whose edges are the orderer pairs (**x**,f(**x**)) for all **x** $\in$ {1,2,...,**n**}
### example of a Functional graph
- **f**: {1,2,...,13} -> {1,2,...,13}
- f(1) = 4
- f(2) = 11
- f(3) = 1
- f(4) = 6
- f(5) = 3
- f(6) = 9
- f(7) = 3
- f(8) = 11
- f(9) = 1
- f(10) = 2
- f(11) = 10
- f(12) = 4
- f(13) = 7
![Graph example](../assets/graph-example.png)
### Fact
- As **n** tends to infinity, the following statements regarding the functional graph of a random function **f** from $F_n$ are true:
1. The expected number of components is $\frac{1}{2}$*ln(**n**)
2. The expected number of points which are on the cycles is $\sqrt{\frac{\pi * n}{2}}$
3. The expected number of **terminal points** (points which have no preimages) is $\frac{n}{e}$
4. The expected number of **k**-th iterate image points (**x** is a **k**-th iterate image point if **x**=f(f(...f(y)...)) for some **y**) is (1-$T_k$)*n, where $T_k$ satisfy the recurrence $T_0$=0,$T_{k+1}$=$e^{-1+T_k}$ for k $\ge$ 0 
### Another definition
- Let **f** be a random function from {1,2,...,**n**} to {1,2,...,**n**} and let **u** $\in$ {1,2,...,**n**}
- Considere the sequence of points $u_0,u_1,u_2,...$ defined by $u_0=u,u_i=f(u_{i-1})$ for $i \ge 1$
- In the graph of f, this sequence describes a path that connects to a cycle
1. The number of edges in the path is called the tail length of u,denoted &lambda;(u)
2. The number of edges in the cycle is called the cycle length of u,denoted &mu;(u)
3. The rho-length of u is the quantity p(u)=&lambda;(u)+&mu;(u)
4. The tree size of u is the number of edges in the maximal tree rooted on a cycle in the component that contains u.
5. The component size of u is the number of edges in the component that contains u.
6. The prodecessors size of u is the number of iterated preimages of u.
### Example
- The functional graph in figure 2.1 has 2 components and 4 terminal points
- The point u=3 has parameters &lambda;(u)=1,&mo;(u)=4,p(u)=5
- The tree,component and predecessors sizes of u=3 are 4,9 and 3
### Fact
- As n tends to infinity
- There are expectatons of some parameters associated with a random point in {1,2,...,n} and a random function from $F_n$
1. Tail length: $\sqrt{\frac{\pi*n}{8}}$
2. Cycle length: $\sqrt{\frac{\pi*n}{8}}$
3. Rho length: $\sqrt{\frac{\pi*n}{2}}$
4. Tree size: $\frac{n}{3}$
5. Component size: $\frac{2n}{3}$
6. predecessors size: $\sqrt{\frac{\pi*n}{8}}$
### Fact
- As n tends to infinity 
- the expectations of the maximum tail,cycle, and rho lengths in a random function from $F_n$ are $c_1$$\sqrt{n}$,$c_2$$\sqrt{n}$ and $c_3$$\sqrt{n}$,respectively,where $c_1$$\approx$ 0,78248, $c_2$$\approx$ 1,73746 and $c_3$$\approx$ 2,4149
## Information theory
### Entropy
- Let X be a random variable which takes a finite set of values $x_1$,$x_2$,...$x_n$ with probability P(X=$x_i$)=$p_i$ where $0 \le p_i \le 1$
- for each i, $1 \le i \le n$ and where $\sum_{i=1}^n$$p_i$=1
- Y and Z are random variables which take on finite sets of values
- The entropy of X is a mathematical measure of the amount of information provided by an observation of X
- It is also the uncertainity about the outcome before an observation of X
- Entropy is also useful for approximating the average number of bits required to encode the elemnts of X
### Definition
- The entropy of uncertainty of X is defined to be H(X)=-$\sum_{i=1}^n$$p_i$$*$lg$*$$p_i$=$\sum_{i=1}^n$$p_1$$*$lg$*$lg($\frac{1}{p_i}$) where by convention $p_i$*lg $p_i$ = $p_i$ * lg ($\frac{1}{p_i}$)=0 if $p_i$=0
### Fact (properties of entropy)
- Let X be a random variable which takes on n values
1. $0 \le H(X) \le lg(n)$
2. H(X)=0 if and only if $p_i$=1 for some i, and $p_j$=0 for all j$\ne$i
3. H(X)=lg(n) if and only if $p_i$=$\frac{1}{n}$ for each i,$1 \le i \le n$ (all the outcomes are equally likely)
### Definition (The joint entropy of X and Y is defined to be)
- H(X,Y)=-$\sum_{x,y}$P(X=x,Y=y)$*$lg(P(X=x,Y=y))
- Where the summation indices x and y range over all values of X and Y,respectively
- The definition can be extended to any number of random variables 
### Fact
- If X and Y are random variables
- $H(X,Y) \le H(X) + H(Y)$, with equality if and only if X and Y are independent
### Definition
- If X,Y are random variables
- The conditional entropy of X and given Y = y is H(X|Y=y)=-$\sum_{x}$$P(X=x|Y=y)$$*$lg(P(X=x|Y=y))
- Where the summation index x ranges over all values of X
- The conditional entropy of X given Y, also called the equivocation of Y about X, is H(X|Y) = $\sum_{y}$P(Y=y)H(X|Y=y), where the summation index y ranges over all values of Y
### Fact (properties of conditional entropy)
- Let X and Y be random variables
1. The quantity H(X|Y) measures the amount of uncertainty remaining about X after Y has been observed
2. H(X,Y)$\ge$ 0 and H(X|X)=0
3. H(X,Y)=H(X)+H(Y|X)=H(Y)+H(X|Y)
4. H(X|Y)$\le$H(X),with equality if and only if X and Y are independent
## Mutual information
### Definition (mutual information or transinformation of random variables X and Y)
- it is I(X;Y)=H(X)-H(X|Y) when the random variables are X and Y
- it is also I(X;Y,Z)=H(X)-H(X|Y,Z) when the random variables are X and the pair Y,Z
### Fact (properties of mutual transformation)
1. The quantity I(X;Y) can be thought of as the amount of information that Y reveals about X
The quantity I(X;Y,Z) can be thought as the amount of information Y and Z reveal about X
2. I(X;Y)$\ge$ 0
3. I(X;Y)=0 if and only if X and Y are independent (that is, Y contributes no information about X)
4. I(X;Y)=I(Y;X)
### Definition (the conditional transformation)
- The conditional transiformation of the pair X,Y given Z is defined to be $I_Z$(X;Y)=H(X|Z)-H(X|Y,Z)
### Fact (properties of conditional transiformation)
1. The quantity $I_Z$(X;Y) can be interpreted as the amount of information that Y provides about X,given that Z has already been observed
2. I(X;Y,Z)=I(X;Y)+$I_Y$(X;Y)
3. $I_Z$(X;Y)=$I_Z$(Y;X)
## Complexity Theory