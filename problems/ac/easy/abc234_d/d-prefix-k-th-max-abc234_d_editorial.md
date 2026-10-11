## D - Prefix K-th Max Editorial

---

The problem statement only asks for the $K$ -th largest value among the first $i$ terms of $P$, but this time, we will consider how to find the set of elements $S_i$ consisting of $K$ largest elements among the first $i$ terms of $P$.

If we naively compute these $N-K+1$ sets, the time complexity will be $O(NK)$ or $O(NK \log K)$, which will exceed the execution time limit. Also, if the $N-K+1$ sets are stored individually in a data structure like an array, the space complexity will be $O(NK)$, which will exceed the execution time limit or the memory limit.

That is why we need the following idea.

- For each $i\ (K \lt i)$, find $S_i$ using the result of $S_{i-1}$. Here, we reuse the data structure used for $S_{i-1}$.

But how can we actually do that?

To tell the conclusion first, we may use the following algorithm.

> First, initialize $S_K$ with $\{P_1,P_2,\ldots,P_K\}$.
> 
> Then, for each $i=K+1,K+2,\ldots,N$ in this order, perform the following operation.
> 
> 1. Initialize $S_i$ with $S_{i-1}$. (Actually we do not do this, because the data structure for the elements to be stored is reused.)
> 2. Let $x$ be the smallest element of $S_{i-1}$. If $x \lt P_i$, then remove $x$ from $S_i$ and insert $P_i$ instead.

The algorithm above still requires $O(NK)$ time when computed naively. However, if we use the data structure called “priority queue” which is provided in most languages, the following operation can be done in $O(\log n)$ time each (where $n$ is the number of elements when the operation is performed), which enables us to improve the time complexity to $O(N \log K)$.

- Retrieve the smallest value
- Remove the smallest value
- Insert an element

What we want to find is the smallest element of $S_i$ for each $i=K,K+1,\ldots,N$. Such element can be of course obtained in a logarithmic time each.

Sample code (Python): In Python, the implementation of “priority queue” is named “heapq”.

```python
import heapqN,K = map(int,input().split())P = list(map(int,input().split()))que = P[0:K]print(min(que))heapq.heapify(que)for i in range(K,N):    minima = heapq.heappop(que)    minima = max(minima,P[i])    heapq.heappush(que,minima)    ans = heapq.heappop(que)    print(ans)    heapq.heappush(que,ans)
```
```
import heapq
N,K = map(int,input().split())
P = list(map(int,input().split()))
que = P[0:K]
print(min(que))
heapq.heapify(que)
for i in range(K,N):
    minima = heapq.heappop(que)
    minima = max(minima,P[i])
    heapq.heappush(que,minima)
    ans = heapq.heappop(que)
    print(ans)
    heapq.heappush(que,ans)
```

Sample code (Python): In C++, the implementation of “priority queue” is named “priority\_queue”. Note that in C++, the element to be popped is by default the largest value, not the smallest.

```
#include<bits/stdc++.h>using namespace std;int main(){    int N,K; cin >> N >> K;    vector<int> P(N);    for(int i=0; i<N; i++) cin >> P[i];    priority_queue<int,vector<int>,greater<int> > que;    for(int i=0; i<K; i++) que.push(P[i]);    cout << que.top() << endl;    for(int i=K; i<N; i++){        if(que.top() < P[i]){            que.pop();            que.push(P[i]);        }        cout << que.top() << endl;    }}
```
```
#include<bits/stdc++.h>
using namespace std;

int main(){
    int N,K; cin >> N >> K;
    vector<int> P(N);
    for(int i=0; i<N; i++) cin >> P[i];
    priority_queue<int,vector<int>,greater<int> > que;
    for(int i=0; i<K; i++) que.push(P[i]);
    cout << que.top() << endl;
    for(int i=K; i<N; i++){
        if(que.top() < P[i]){
            que.pop();
            que.push(P[i]);
        }
        cout << que.top() << endl;
    }
}
```

posted:  
last update:

---
