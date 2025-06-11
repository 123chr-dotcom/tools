# 编译器设置
CXX = g++
MOC = moc
UIC = uic

# 编译选项
CXXFLAGS = -std=c++17 -Wall -fPIC
QT_CFLAGS = $(shell pkg-config --cflags Qt5Widgets)
QT_LDFLAGS = $(shell pkg-config --libs Qt5Widgets)

# 使用pkg-config获取完整Qt链接参数
QT_LIBS = $(shell pkg-config --libs Qt5Widgets Qt5Gui Qt5Core)

# 最终编译标志
ALL_CFLAGS = $(CXXFLAGS) $(shell pkg-config --cflags Qt5Widgets Qt5Gui Qt5Core)
ALL_LDFLAGS = $(QT_LIBS) -Wl,-subsystem,windows

# moc预处理
MOC_FILES = moc_MainWindow.cpp

# 源文件
SOURCES = main.cpp MainWindow.cpp $(MOC_FILES)
HEADERS = MainWindow.h
UI = MainWindow.ui

# 生成的目标文件
OBJECTS = $(SOURCES:.cpp=.o)
TARGET = app

# UI生成的文件
UI_HEADERS = ui_MainWindow.h

# 默认目标
all: $(UI_HEADERS) $(MOC_FILES) $(TARGET)

# 主目标
$(TARGET): $(OBJECTS)
	$(CXX) -o $@ $(OBJECTS) $(ALL_LDFLAGS)

# 从.ui文件生成头文件
ui_%.h: %.ui
	$(UIC) -o $@ $<

# moc处理
moc_%.cpp: %.h
	$(MOC) $< -o $@

# 编译.cpp文件
%.o: %.cpp
	$(CXX) $(ALL_CFLAGS) -c -o $@ $<

# 从.ui文件生成头文件
$(UI_HEADERS): $(UI)
	$(UIC) -o $@ $<
	
# moc处理
$(MOC_FILES): $(HEADERS)
	$(MOC) $< -o $@

# 清理
clean:
	rm -f $(TARGET) $(OBJECTS) $(MOC_FILES) ui_*.h

.PHONY: all clean
