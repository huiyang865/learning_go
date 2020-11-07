def main():
    print('hello world!')  # 默认换行

    print('hello', end='')  # 不换行
    print(' world!')

    # 变量的定义
    i = 40
    print(i)

    i = 'hello world!'
    print(i)

    # 定义匿名变量
    my_list = range(10)
    for _, item_num in enumerate(my_list):
        print(item_num)

    # 常量定义，python中常量是可以被修改
    PI = 3.1415926
    r = 20
    c = 2 * PI * r
    print(c)


if __name__ == "__main__":
    main()  # 需要手动调用
