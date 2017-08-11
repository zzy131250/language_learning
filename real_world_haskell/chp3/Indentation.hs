foo = let firstDefinition = blah blah
          -- 只有注释的行被视作空白行
              continuation blah
          -- 减少缩进，于是下面这行就变成了一行新定义
          secondDefinition = yada yada
              continuation yada
      in whatever
