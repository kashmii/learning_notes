module Logggg
  private # ここがあるかないかで変わる

  def log(text)
    puts "[LOG] #{text}"
  end
end

class Usssss
  include Logggg

  def name
    log 'name is called'
    'Ring ring!!'
  end
end

ussuss = Usssss.new
ussuss.name

# privateだと以下がエラーになる
ussuss.log 'p?????'

# ===
# ===

# ポイント：
#   ミックスイン ... クラスにモジュールを（大抵includeで）取り込むこと