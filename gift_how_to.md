# How to download and use eth-tool
## Prerequisites
- First you will need to download and install **git** (https://git-scm.com/downloads)
  - **git** is a tool for managing **source code**
  - you need it because you will be downloading my source code and building the program from scratch
- Next you will need to create an account on **github** (https://github.com) so I can share access to the source code
  - Text me your **github username** and I will grant access
-  Next you will need download and install **Go** (https://go.dev/)
  - **Go** is the programming language used to write  the tool
## Downloading the source code with git
- If you are on **Windows** open up **Git Bash** which should have been installed automatically when you installed git
- If you are on **Mac** open up **Terminal**
- Type each of these commands ending with hitting **enter** on your keyboard:
	- mkdir ~/repos && cd repos
	- git clone https://github.com/donahoem/eth-tool.git
	- cd eth-tool
	- go get
	- go build .
- That's it!  The tool is build and ready for use
## Using the tool
- If you still have **Terminal** or **Git Bash** open from the previous step, good.
- If not, open it up and type:
	- cd ~/repos/eth-tool
- Now you can run the tool by typing:
	- ./eth-tool
- The above command does not tell the tool to do anything, but prints it's **usage**
- To tell the tool to **create a new wallet**, run the following command:
	- ./eth-tool --create-wallet
	- You should now see the address and private key on the screen
- To **check the balance** of this address (or any address), run the following command, replacing \<address\> with the address you're interested in:
	- ./eth-tool --check-balance \<address\>
	- You should see the amount of **wei** in the wallet on the screen
	- *NOTE:* 1 _ETH_ = 1000000000000000000 _Wei_
- To **send ethereum** from one address to antoehr, run the following command:
	- ./eth-tool --send-wei \<sender private key\> <amount in wei\> \<recipient address\> 
	- Replace \<sender private key\> with the private key of the address you wish to send from
	- Replace \<amount in wei\> with the amount of wei you wish to send
	- Replace \<recipient address\> with the address that you wish to send to
