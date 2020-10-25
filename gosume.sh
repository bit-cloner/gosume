#! /bin/bash
#Check for required binaries : saml2aws and assume
B1=./saml2aws
B2=./assume
F1=~/.saml2aws
if test -f "$B1"; then
    echo "$B1 exists."
else
    # wget or curl binary from github 
    wget https://github.com/bit-cloner/gosume/releases/download/0.9/saml2aws
    chmod +x ./saml2aws
fi
if test -f "$B2"; then
    echo "$B2 exists."
else
    # wget or curl binary from github 
    wget https://github.com/bit-cloner/gosume/releases/download/0.9/assume
    chmod +x ./assume
fi
# Login to Identity Provider to create a credential profile
# Check if .saml2aws file exists 
if test -f "$F1"; then
    echo "$F1 exists. Identity Provider has been configured"
else
    ./saml2aws configure
fi 
~/saml2aws login
echo -e "You have been authenticated with AWS app configured in Identity provider.\n you can now assume a role"
echo -e "Enter the arn of the role to be assumed"
read arn
./assume ${arn}


