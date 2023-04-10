在Mac上设置两个Git账号的过程与在其他操作系统上设置基本相同。您可以按照以下步骤进行设置：

生成SSH密钥对
使用以下命令生成SSH密钥对：

ssh-keygen -t rsa -C "your-email-address"
将“your-email-address”替换为您Git账户的注册邮箱。在提示输入文件名和密码时，可以保留默认值并按Enter键。

添加SSH密钥到Git账户
登录到Git网站，例如GitHub，转到设置（Settings）菜单，然后选择SSH Keys选项卡。点击"Add SSH Key"按钮，在“key”字段中粘贴您刚才生成的公钥(~/.ssh/id_rsa.pub)内容，并为该密钥提供一个名称。

配置本地Git
在您的Mac上打开终端窗口，进入您要使用的Git项目的目录，然后运行以下命令：

git config user.name "Your Name"
git config user.email "your-email-address"
将“Your Name”替换为您希望显示在Git提交历史记录中的名称，将“your-email-address”替换为您Git账户的注册邮箱。

配置SSH主机别名
如果您需要同时访问多个Git服务器，您可以配置每个服务器的别名。编辑~/.ssh/config文件，添加以下条目：

Host alias1
    HostName gitserver1.com
    User git
    IdentityFile ~/.ssh/id_rsa

Host alias2
    HostName gitserver2.com
    User git
    IdentityFile ~/.ssh/id_rsa_other
将“alias1”和“alias2”替换为您希望使用的别名，将“gitserver1.com”和“gitserver2.com”替换为服务器的地址。如果您生成了多个SSH密钥对，请确保在每个主机别名中使用正确的IdentityFile路径。

现在，您可以通过运行“git clone alias1:repository.git”或“git clone alias2:repository.git”等命令来克隆不同的Git存储库并使用不同的Git账户。