import time
import random
import csv


def generate(n):
	p,w,d = [],[],[]
	for i in range(n):
		p.append(random.randint(1,100))
		w.append(random.randint(1,10))
	ep = sum(p)
	minW = int(0.2*ep)
	maxW = int(0.6*ep)
	for i in range(n):
		d.append(random.randint(minW, maxW))
	return p,w,d


def randomPath(n):
    seq=[]
    for i in range(n):
        seq.append(random.randint(1,n));
        while seq[i] in seq[0:i]:
            seq[i]=random.randint(1,n);
    return seq


n=10
pathdetail=[]
path =[]

for i in range(n):
    path.append(i+1)
    pathdetail.append([generate(n)[0][i], generate(n)[1][i], generate(n)[2][i]])
path=randomPath(n)

with open("path_genetics.csv", "w") as new_file:
    csv_writer = csv.writer(new_file, delimiter=' ')
    csv_writer.writerow(['Node', 'Start', 'End', 'Distance'])
    for i in range(1, n + 1):
        csv_writer.writerow([i, pathdetail[i-1][0],pathdetail[i-1][1], pathdetail[i-1][2]])

def cost(seq):
    ct,ci = 0, 0
    for i in range(len(seq)):
        ci += pathdetail[seq[i]-1][0]
        ct += pathdetail[seq[i]-1][1] * max(0, ci - pathdetail[seq[i]-1][2])
    return ct


def permutation(s, i, j):
    seq=list(s)
    seq[i], seq[j] = seq[j], seq[i]
    return seq
def insertion(s, i, j):
    temp=s[j]
    del s[j]
    s.insert(i,temp)
    seq = s
    return seq
def left_pivot(l,i):
    l1 = l[0:i]
    l1.reverse()
    l2 = l[i:len(l)]
    l = l1+l2
    return l

def optimization(seq_arg):
    i = 0
    timeout = time.time()
    xp = []
    cost_end = 99999999999999999999999
    seq_end = []
    while True:
        if time.time() > timeout:
            break
        if i == 0:
            min_i = random.randint(0, len(seq_arg) - 1)
            min_j = random.randint(0, len(seq_arg) - 1)
            while min_i == min_j:
                min_i = random.randint(0,len(seq_arg)-1)
            xp = permutation(seq_arg, min_i,min_j)
            cost_xp = cost(xp)
            cost_min_xseconde = cost(permutation(xp, 0, 1))
            for k in range(len(xp)):
                for j in range(k+1, len(xp)):
                    xseconde = permutation(xp, k, j)
                    cost_xsec = cost(xseconde)
                    if cost(xseconde) < cost_min_xseconde:
                        xseconde_min = xseconde
                        cost_min_xseconde = cost_xsec
            if cost_min_xseconde < cost_xp:
                xp = xseconde_min
                i = 0
                seq_arg = xseconde_min
            else:
                seq_arg = xp
                i = 1
        if i == 1:
            min_i = random.randint(0, len(seq_arg) - 1)
            min_j = random.randint(0, len(seq_arg) - 1)
            while min_i == min_j:
                min_i = random.randint(0, len(seq_arg) - 1)
            xp = insertion(seq_arg, min_i,min_j)
            cout_xp = cost(xp)
            cout_min_xseconde = cost(insertion(xp,0,2))
            for k in range(len(xp)):
                for j in range(k+1,len(xp)):
                    xseconde = insertion(xp, k, j)
                    cout_xsec = cost(xseconde)
                    if cost(xseconde) < cout_min_xseconde:
                        xseconde_min = xseconde
                        cout_min_xseconde = cout_xsec
            if cout_min_xseconde < cout_xp:
                xp = xseconde_min
                i=0
                seq_arg = xseconde_min
            else:
                seq_arg = xp
                i=2
        if i == 2:
            xp = left_pivot(seq_arg, random.randint(0,len(seq_arg)-1))
            cout_xp = cost(xp)
            cout_min_xseconde = cost(left_pivot(xp,2))
            for k in range(2,len(seq_arg)):
                xseconde = left_pivot(xp, k)
                cout_xsec = cost(xseconde)
                if cost(xseconde) < cout_min_xseconde:
                    xseconde_min = xseconde
                    cout_min_xseconde = cout_xsec
            if cout_min_xseconde < cout_xp:
                xp = xseconde_min
                seq_arg = xseconde_min
            i = 0
        if cost_min_xseconde < cost_end:
            cost_end = cost_min_xseconde
            seq_end = xseconde_min
    return seq_end, cost_end


new_path = []
for i in range(len(path)-1):
    new_path.append(path[i])
min_seq, min_cost = optimization(new_path)
min_seq.append(path[n-1])
min_seq.insert(0, path[0])

result=f"Path: {path}\nDistance: {cost(path)}\nNew Path: {min_seq} \nNew Distance: {min_cost}"
with open("results_genetics.csv", "w") as new_file:
    new_file.write(result)