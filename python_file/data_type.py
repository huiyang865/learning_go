import sys


def main():
    # 定义一个整数类型，显示内存占用大小
    # 对于python，内存占用大小是动态的，根据数据大小动态分配，不需要担心上限
    age = 18
    print(sys.getsizeof(age))  # 占用28个字节
    print(type(71.2))


if __name__ == "__main__":
    main()