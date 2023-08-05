using System.Security.Policy;
using System.Threading.Tasks;
using System.Windows.Forms;
using Microsoft.Web.WebView2.Core;
using Microsoft.Web.WebView2.WinForms;


namespace PosidoniaSaver
{
    public partial class Form1 : Form
    {
        String author = "Yiðit GÜMÜÞ | <github.com/yigitoo>";
        public Form1()
        {
            Console.WriteLine($"Maded by {author}");
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            InitBrowser("http://localhost:1234/");
        }

        private async Task initizated()
        {
            await webView21.EnsureCoreWebView2Async(null);
        }

        private async void InitBrowser(String url)
        {
            await initizated();
            webView21.CoreWebView2.Navigate(url);

        }

        void WebView_ServerCertificateErrorDetected(object sender, CoreWebView2ServerCertificateErrorDetectedEventArgs e)
        {
            CoreWebView2Certificate certificate = e.ServerCertificate;
            e.Action = CoreWebView2ServerCertificateErrorAction.AlwaysAllow;
        }

    }
}
