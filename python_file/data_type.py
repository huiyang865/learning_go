import sys


def main():
    # 定义一个整数类型，显示内存占用大小
    # 对于python，内存占用大小是动态的，根据数据大小动态分配，不需要担心上限
    age = 18
    print(sys.getsizeof(age))  # 占用28个字节
    print(type(71.2))

    # python的类型比go少很多int float
    # 1. 字符串转int
    data = int('21')
    print(type(data), data)
    data = int('21', 16)  # 设置输入为16进制数值
    print(type(data), data)  # 输出为33

    # 2. int转字符串
    data_str = str(int(data))
    print(type(data_str), data_str)

    # 3. float转字符串
    data_str = str(3.14)
    print(type(data_str), data_str)

    # 4. bool类型转换
    bool_data = bool('false')  # 只要字符串非空，转换都为True，空字符串转为False
    print(type(bool_data), bool_data)
    # 应用strtobool
    from distutils.util import strtobool
    bool_data = strtobool('false')
    print(type(bool_data), bool_data)
    bool_data = strtobool('False')
    print(type(bool_data), bool_data)
    bool_data = strtobool('1')
    print(type(bool_data), bool_data)
    bool_data = strtobool('TRUE')
    print(type(bool_data), bool_data)
    # bool_data = strtobool('Trues') 会报错，只能转false, true, 1, 0字符串


if __name__ == "__main__":
    main()