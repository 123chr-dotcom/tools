#include "MainWindow.h"
#include <QPushButton>
#include <QWidget>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    setupUI();
}

MainWindow::~MainWindow()
{
}

void MainWindow::setupUI()
{
    // 创建一个简单的按钮
    QPushButton *button = new QPushButton("点击我", this);
    button->setGeometry(50, 50, 100, 30);
    
    // 设置窗口大小
    setFixedSize(200, 150);
    setWindowTitle("Qt示例程序");
}
