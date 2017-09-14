using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace CalculadoraRPC
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
        }

        private void CalculatorButton_Click(object sender, RoutedEventArgs e)
        {
            Button btn_Sender = sender as Button;
            tbk_DisplayContent.Text = String.Format("{0}{1}", tbk_DisplayContent.Text, (btn_Sender.Content as String));
        }

        private void CalculatorClearButton_Click(object sender, RoutedEventArgs e)
        {
            Button btn_Sender = sender as Button;
            tbk_DisplayContent.Text = String.Empty;
        }

        private void CalculatorBackspaceButton_Click(object sender, RoutedEventArgs e)
        {
            Button btn_Sender = sender as Button;
            if (tbk_DisplayContent.Text.Length > 0)
                tbk_DisplayContent.Text = tbk_DisplayContent.Text.Remove(tbk_DisplayContent.Text.Length - 1);
        }

        private async void CalculatorFunctionButton_Click(object sender, RoutedEventArgs e)
        {
            Button btn_Sender = sender as Button;
            tbk_DisplayContent.Text = await CalcController.CalculateFunction(btn_Sender.Content as String, tbk_DisplayContent.Text);
        }

        private async void CalculatorEqualButton_Click(object sender, RoutedEventArgs e)
        {
            tbk_DisplayContent.Text = await CalcController.CalculateExpression(tbk_DisplayContent.Text);
        }
    }
}
