import sys

class Defer:
    '''
    模拟 golang 的 defer
    python 的 try...finally... 和 with 语句都能实现函数结尾保证做一件事情，但是没有 golang 的 defer 那么
    优雅，比如说会增加缩进，这里参考 C++ 的 RAII 技术实现一个类似 golang 的 defer 的功能：构造函数传入 finally
    要执行的函数极其执行时要调用的参数
    '''
    def __init__(self, fn, *args):
        print(">>> new Defer object")
        self.defer_func, self.defer_argv = fn, args

    def __del__(self):
        print(">>> call Defer.__del__()")
        self.defer_func(*self.defer_argv)


def catFile(fn):
    '''
    这是一个使用 Defer 类的 demo, 函数打开一个文件，然后打印文件的内容，最后保证在函数执行完成的时候会调用
    fd.close() 关闭文件
    '''
    fd = open(fn, encoding="utf-8")
    defer = Defer(fd.close)  # variable defer will be deleted before return

    print("".join(fd.readlines()), end="")


catFile(sys.argv[1])
