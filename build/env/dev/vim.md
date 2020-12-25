vi/vim 的使用
基本上 vi/vim 共分为三种模式，分别是命令模式（Command mode），输入模式（Insert mode）和底线命令模式（Last line mode）。 这三种模式的作用分别是：

命令模式：
用户刚刚启动 vi/vim，便进入了命令模式。

此状态下敲击键盘动作会被Vim识别为命令，而非输入字符。比如我们此时按下i，并不会输入一个字符，i被当作了一个命令。

以下是常用的几个命令：

i 切换到输入模式，以输入字符。
x 删除当前光标所在处的字符。
: 切换到底线命令模式，以在最底一行输入命令。
若想要编辑文本：启动Vim，进入了命令模式，按下i，切换到输入模式。

命令模式只有一些最基本的命令，因此仍要依靠底线命令模式输入更多命令。

输入模式
在命令模式下按下i就进入了输入模式。

在输入模式中，可以使用以下按键：

字符按键以及Shift组合，输入字符
ENTER，回车键，换行
BACK SPACE，退格键，删除光标前一个字符
DEL，删除键，删除光标后一个字符
方向键，在文本中移动光标
HOME/END，移动光标到行首/行尾
Page Up/Page Down，上/下翻页
Insert，切换光标为输入/替换模式，光标将变成竖线/下划线
ESC，退出输入模式，切换到命令模式
底线命令模式
在命令模式下按下:（英文冒号）就进入了底线命令模式。

底线命令模式可以输入单个或多个字符的命令，可用的命令非常多。

在底线命令模式中，基本的命令有（已经省略了冒号）：

q 退出程序
w 保存文件
按ESC键可随时退出底线命令模式。

vi/vim 使用实例
使用 vi/vim 进入一般模式
如果你想要使用 vi 来建立一个名为 runoob.txt 的文件时，你可以这样做：
```sh
$ vim runoob.txt
```
按下 i 进入输入模式(也称为编辑模式)，开始编辑文字
在一般模式之中，只要按下 i, o, a 等字符就可以进入输入模式了！

在编辑模式当中，你可以发现在左下角状态栏中会出现 –INSERT- 的字样，那就是可以输入任意字符的提示。

这个时候，键盘上除了 Esc 这个按键之外，其他的按键都可以视作为一般的输入按钮了，所以你可以进行任何的编辑。

按下 ESC 按钮回到一般模式
好了，假设我已经按照上面的样式给他编辑完毕了，那么应该要如何退出呢？是的！没错！就是给他按下 Esc 这个按钮即可！马上你就会发现画面左下角的 – INSERT – 不见了！

在一般模式中按下 :wq 储存后离开 vi
OK，我们要存档了，存盘并离开的指令很简单，输入 :wq 即可保存离开！

第一部份：一般模式可用的按钮说明，游标移动、复制贴上、搜寻取代等

移动游标的方法
h 或 向左方向键(←)	游标向左移动一个字元
j 或 向下方向键(↓)	游标向下移动一个字元
k 或 向上方向键(↑)	游标向上移动一个字元
l 或 向右方向键(→)	游标向右移动一个字元
如果你将右手放在键盘上的话，你会发现 hjkl 是排列在一起的，因此可以使用这四个按钮来移动游标。 如果想要进行多次移动的话，例如向下移动 30 行，可以使用 "30j" 或 "30↓" 的组合按键， 亦即加上想要进行的次数(数字)后，按下动作即可！
[Ctrl] + [f]	荧幕‘向下’移动一页，相当于 [Page Down]按键 (常用)
[Ctrl] + [b]	荧幕‘向上’移动一页，相当于 [Page Up] 按键 (常用)
[Ctrl] + [d]	荧幕‘向下’移动半页
[Ctrl] + [u]	荧幕‘向上’移动半页
+	游标移动到非空白字元的下一列
-	游标移动到非空白字元的上一列
n<space>	那个 n 表示‘数字’，例如 20 。按下数字后再按空白键，游标会向右移动这一行的 n 个字元。例如 20<space> 则游标会向后面移动 20 个字元距离。
0 或功能键[Home]	这是数字‘ 0 ’：移动到这一行的最前面字元处 (常用)
$ 或功能键[End]	移动到这一行的最后面字元处(常用)
H	游标移动到这个荧幕的最上方那一行的第一个字元
M	游标移动到这个荧幕的中央那一行的第一个字元
L	游标移动到这个荧幕的最下方那一行的第一个字元
G	移动到这个档案的最后一行(常用)
nG	n 为数字。移动到这个档案的第 n 行。例如 20G 则会移动到这个档案的第 20 行(可配合 :set nu)
gg	移动到这个档案的第一行，相当于 1G 啊！ (常用)
n<Enter>	n 为数字。游标向下移动 n 行(常用)
搜寻与取代
/word	向游标之下寻找一个名称为 word 的字串。例如要在档案内搜寻 vbird 这个字串，就输入 /vbird 即可！ (常用)
?word	向游标之上寻找一个字串名称为 word 的字串。
n	这个 n 是英文按键。代表‘重复前一个搜寻的动作’。举例来说， 如果刚刚我们执行 /vbird 去向下搜寻 vbird 这个字串，则按下 n 后，会向下继续搜寻下一个名称为 vbird 的字串。如果是执行 ?vbird 的话，那么按下 n 则会向上继续搜寻名称为 vbird 的字串！
N	这个 N 是英文按键。与 n 刚好相反，为‘反向’进行前一个搜寻动作。 例如 /vbird 后，按下 N 则表示‘向上’搜寻 vbird 。
使用 /word 配合 n 及 N 是非常有帮助的！可以让你重复的找到一些你搜寻的关键字！
:n1,n2s/word1/word2/g	n1 与 n2 为数字。在第 n1 与 n2 行之间寻找 word1 这个字串，并将该字串取代为 word2 ！举例来说，在 100 到 200 行之间搜寻 vbird 并取代为 VBIRD 则：
‘:100,200s/vbird/VBIRD/g’。(常用)
:1,$s/word1/word2/g	从第一行到最后一行寻找 word1 字串，并将该字串取代为 word2 ！(常用)
:1,$s/word1/word2/gc	从第一行到最后一行寻找 word1 字串，并将该字串取代为 word2 ！且在取代前显示提示字元给使用者确认 (confirm) 是否需要取代！(常用)
删除、复制与贴上
x, X	在一行字当中，x 为向后删除一个字元 (相当于 [del] 按键)， X 为向前删除一个字元(相当于 [backspace] 亦即是倒退键) (常用)
nx	n 为数字，连续向后删除 n 个字元。举例来说，我要连续删除 10 个字元， ‘10x’。
dd	删除游标所在的那一整列(常用)
ndd	n 为数字。删除游标所在的向下 n 列，例如 20dd 则是删除 20 列 (常用)
d1G	删除游标所在到第一行的所有资料
dG	删除游标所在到最后一行的所有资料
d$	删除游标所在处，到该行的最后一个字元
d0	那个是数字的 0 ，删除游标所在处，到该行的最前面一个字元
yy	复制游标所在的那一行(常用)
nyy	n 为数字。复制游标所在的向下 n 列，例如 20yy 则是复制 20 列(常用)
y1G	复制游标所在列到第一列的所有资料
yG	复制游标所在列到最后一列的所有资料
y0	复制游标所在的那个字元到该行行首的所有资料
y$	复制游标所在的那个字元到该行行尾的所有资料
p, P	p 为将已复制的资料在游标下一行贴上，P 则为贴在游标上一行！ 举例来说，我目前游标在第 20 行，且已经复制了 10 行资料。则按下 p 后， 那 10 行资料会贴在原本的 20 行之后，亦即由 21 行开始贴。但如果是按下 P 呢？ 那么原本的第 20 行会被推到变成 30 行。 (常用)
J	将游标所在列与下一列的资料结合成同一列
c	重复删除多个资料，例如向下删除 10 行，[ 10cj ]
u	复原前一个动作。(常用)
[Ctrl]+r	重做上一个动作。(常用)
这个 u 与 [Ctrl]+r 是很常用的指令！一个是复原，另一个则是重做一次～ 利用这两个功能按键，你的编辑，嘿嘿！很快乐的啦！
.	不要怀疑！这就是小数点！意思是重复前一个动作的意思。 如果你想要重复删除、重复贴上等等动作，按下小数点‘.’就好了！ (常用)

第二部份：一般模式切换到编辑模式的可用的按钮说明

进入插入或取代的编辑模式
i, I	进入插入模式(Insert mode)：
i 为‘从目前游标所在处插入’， I 为‘在目前所在行的第一个非空白字元处开始插入’。 (常用)
a, A	进入插入模式(Insert mode)：
a 为‘从目前游标所在的下一个字元处开始插入’， A 为‘从游标所在行的最后一个字元处开始插入’。(常用)
o, O	进入插入模式(Insert mode)：
这是英文字母 o 的大小写。o 为‘在目前游标所在的下一行处插入新的一行’； O 为在目前游标所在处的上一行插入新的一行！(常用)
r, R	进入取代模式(Replace mode)：
r 只会取代游标所在的那一个字元一次；R会一直取代游标所在的文字，直到按下 ESC 为止；(常用)
上面这些按键中，在 vi 画面的左下角处会出现‘--INSERT--’或‘--REPLACE--’的字样。 由名称就知道该动作了吧！！特别注意的是，我们上面也提过了，你想要在档案里面输入字元时， 一定要在左下角处看到 INSERT 或 REPLACE 才能输入喔！
[Esc]	退出编辑模式，回到一般模式中(常用)

第三部份：一般模式切换到指令列模式的可用的按钮说明

指令列的储存、离开等指令
:w	将编辑的资料写入硬碟档案中(常用)
:w!	若档案属性为‘唯读’时，强制写入该档案。不过，到底能不能写入， 还是跟你对该档案的档案权限有关啊！
:q	离开 vi (常用)
:q!	若曾修改过档案，又不想储存，使用 ! 为强制离开不储存档案。
注意一下啊，那个惊叹号 (!) 在 vi 当中，常常具有‘强制’的意思～
:wq	储存后离开，若为 :wq! 则为强制储存后离开 (常用)
ZZ	这是大写的 Z 喔！若档案没有更动，则不储存离开，若档案已经被更动过，则储存后离开！
:w [filename]	将编辑的资料储存成另一个档案（类似另存新档）
:r [filename]	在编辑的资料中，读入另一个档案的资料。亦即将 ‘filename’ 这个档案内容加到游标所在行后面
:n1,n2 w [filename]	将 n1 到 n2 的内容储存成 filename 这个档案。
:! command	暂时离开 vi 到指令列模式下执行 command 的显示结果！例如
‘:! ls /home’即可在 vi 当中察看 /home 底下以 ls 输出的档案资讯！
vim 环境的变更
:set nu	显示行号，设定之后，会在每一行的字首显示该行的行号
:set nonu	与 set nu 相反，为取消行号！

特别注意，在 vi 中，‘数字’是很有意义的！数字通常代表重复做几次的意思！ 也有可能是代表去到第几个什么什么的意思。举例来说，要删除 50 行，则是用 ‘50dd’ 对吧！ 数字加在动作之前～那我要向下移动 20 行呢？那就是‘20j’或者是‘20↓’即可。

OK！会这些指令就已经很厉害了，因为常用到的指令也只有不到一半！通常 vi 的指令除了上面鸟哥注明的常用的几个外，其他是不用背的，你可以做一张简单的指令表在你的荧幕墙上， 一有疑问可以马上的查询呦！这也是当初鸟哥使用 vim 的方法啦！

区块选择的按键意义
v	字元选择，会将游标经过的地方反白选择！
V	行选择，会将游标经过的行反白选择！
[Ctrl]+v	区块选择，可以用长方形的方式选择资料
y	将反白的地方复制起来
d	将反白的地方删除掉

多档案编辑的按键
:n	编辑下一个档案
:N	编辑上一个档案
:files	列出目前这个 vim 的开启的所有档案

当我有一个档案非常的大，我查阅到后面的资料时，想要‘对照’前面的资料， 是否需要使用 [ctrl]+f 与 [ctrl]+b (或 pageup, pagedown 功能键) 来跑前跑后查阅？

我有两个需要对照着看的档案，不想使用前一小节提到的多档案编辑功能；

多视窗情况下的按键功能
:sp [filename]	开启一个新视窗，如果有加 filename， 表示在新视窗开启一个新档案，否则表示两个视窗为同一个档案内容(同步显示)。
[ctrl]+w+ j
[ctrl]+w+↓	按键的按法是：先按下 [ctrl] 不放， 再按下 w 后放开所有的按键，然后再按下 j (或向下方向键)，则游标可移动到下方的视窗。
[ctrl]+w+ k
[ctrl]+w+↑	同上，不过游标移动到上面的视窗。
[ctrl]+w+ q	其实就是 :q 结束离开啦！ 举例来说，如果我想要结束下方的视窗，那么利用 [ctrl]+w+↓ 移动到下方视窗后，按下 :q 即可离开， 也可以按下 [ctrl]+w+q 啊！

vim 的环境设定参数
:set nu
:set nonu	就是设定与取消行号啊！
:set hlsearch
:set nohlsearch	hlsearch 就是 high light search(高亮度搜寻)。 这个就是设定是否将搜寻的字串反白的设定值。预设值是 hlsearch
:set autoindent
:set noautoindent	是否自动缩排？autoindent 就是自动缩排。
:set backup	是否自动储存备份档？一般是 nobackup 的， 如果设定 backup 的话，那么当你更动任何一个档案时，则原始档案会被另存成一个档名为 filename~ 的档案。 举例来说，我们编辑 hosts ，设定 :set backup ，那么当更动 hosts 时，在同目录下，就会产生 hosts~ 档名的档案，记录原始的 hosts 档案内容
:set ruler	还记得我们提到的右下角的一些状态列说明吗？ 这个 ruler 就是在显示或不显示该设定值的啦！
:set showmode	这个则是，是否要显示 --INSERT-- 之类的字眼在左下角的状态列。
:set backspace=(012)	一般来说， 如果我们按下 i 进入编辑模式后，可以利用倒退键 (backspace) 来删除任意字元的。 但是，某些 distribution 则不许如此。此时，我们就可以透过 backspace 来设定啰～ 当 backspace 为 2 时，就是可以删除任意值；0 或 1 时，仅可删除刚刚输入的字元， 而无法删除原本就已经存在的文字了！
:set all	显示目前所有的环境参数设定值。
:set	显示与系统预设值不同的设定参数， 一般来说就是你有自行变动过的设定参数啦！
:syntax on
:syntax off	是否依据程式相关语法显示不同颜色？ 举例来说，在编辑一个纯文字档时，如果开头是以 # 开始，那么该行就会变成蓝色。 如果你懂得写程式，那么这个 :syntax on 还会主动的帮你除错呢！但是， 如果你仅是编写纯文字档案，要避免颜色对你的荧幕产生的干扰，则可以取消这个设定 。
:set bg=dark
:set bg=light	可用以显示不同的颜色色调，预设是‘ light ’。如果你常常发现注解的字体深蓝色实在很不容易看， 那么这里可以设定为 dark 喔！试看看，会有不同的样式呢！

总之，这些设定值很有用处的啦！但是......我是否每次使用 vim 都要重新设定一次各个参数值？ 这不太合理吧？没错啊！所以，我们可以透过设定档来直接规定我们习惯的 vim 操作环境呢！ 整体 vim 的设定值一般是放置在 /etc/vimrc 这个档案，不过，不建议你修改他！ 你可以修改 ~/.vimrc 这个档案 (预设不存在，请你自行手动建立！)，将你所希望的设定值写入！ 举例来说，可以是这样的一个档案：