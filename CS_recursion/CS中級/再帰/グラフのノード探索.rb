# グラフクラスの定義
class Graph
  def initialize(n)
    @graph = Array.new(n) { [] }
  end

  def add_edge(a, b)
    @graph[a] << b
  end

  def [](index)
    @graph[index]
  end
end

# ノード v を探索し、v から 1 ステップで行くことのできるノードたちを再帰的に探索する
def rec(v, graph, seen)
  seen[v] = true
  graph[v].each do |next_node|
    rec(next_node, graph, seen) unless seen[next_node] # 既に訪問済みなら探索しない
  end
end

# メインの処理
n, m = gets.split.map(&:to_i) # 頂点数と枝数
graph = Graph.new(n) # グラフの初期化

m.times do
  a, b = gets.split.map(&:to_i) # ノード a からノード b へと有向辺を張る
  graph.add_edge(a, b)
end

# 探索
seen = Array.new(n, false) # 初期状態では全ノードが未訪問
(0...n).each do |v|
  next if seen[v] # 既に訪問済みなら探索しない
  rec(v, graph, seen)
end
