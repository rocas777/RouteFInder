import time
import random

path=['walk_1143317559', 'walk_1143317946', 'walk_1143316551', 'walk_2902711664']
pathdetail=[['walk_1143317559','walk_1143317946',88.50696628450761],
            ['walk_1143317946','walk_1143316551',20.93599453900186],
            ['walk_1143316551','walk_2902711664',6.8670632025406295]]

def cost(seq):
    ct,ci = 0, 0
    for i in range(len(pathdetail)):
        ct += pathdetail[i][2]
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


def optimisation(seq_arg):
    i = 0
    delais_desire = 0
    timeout = time.time() + delais_desire
    xp = []
    cost_fin = 99999999999999999999999
    seq_fin = []
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
            min_j = random.randint(1, len(seq_arg) - 1)
            while min_i == min_j:
                min_i = random.randint(0, len(seq_arg) - 1)
            xp = insertion(seq_arg, min_i,min_j)
            cost_xp = cost(xp)
            cost_min_xseconde = cost(insertion(xp,0,2))
            for k in range(len(xp)):
                for j in range(k+1,len(xp)):
                    xseconde = insertion(xp, k, j)
                    cost_xsec = cost(xseconde)
                    if cost(xseconde) < cost_min_xseconde:
                        xseconde_min = xseconde
                        cost_min_xseconde = cost_xsec
            if cost_min_xseconde < cost_xp:
                xp = xseconde_min
                i=0
                seq_arg = xseconde_min
            else:
                seq_arg = xp
                i=2
        if i == 2:
            xp = left_pivot(seq_arg, random.randint(0,len(seq_arg)-1))
            cost_xp = cost(xp)
            cost_min_xseconde = cost(left_pivot(xp,2))
            for k in range(2,len(seq_arg)):
                xseconde = left_pivot(xp, k)
                cost_xsec = cost(xseconde)
                if cost(xseconde) < cost_min_xseconde:
                    xseconde_min = xseconde
                    cost_min_xseconde = cost_xsec
            if cost_min_xseconde < cost_xp:
                xp = xseconde_min
                seq_arg = xseconde_min
            i = 0
        if cost_min_xseconde < cost_fin:
            cost_fin = cost_min_xseconde
    return seq_fin, cost_fin


new_path, distance = optimisation(path)
injection_resultat=f"Path: {path}\nOptimized Path: {new_path}\nDistance: [{distance}]"
with open("results_genetics.csv", "w") as new_file:
    new_file.write(injection_resultat)