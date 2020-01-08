class Node:
  def __init__(self, value, node):
      self.value = value
      self.node = node

b = Node(50, None)
n = Node(5, b)

def addElement(n,value):
    if n.node == None:
       v = Node(value, None)
       n.node = v
       return n
    return addElement(n.node, value)

addElement(n, 6)
print(n.node.value, n.value)

