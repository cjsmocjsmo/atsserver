#!/bin/sh
addr="http://192.168.0.91:8080"
# addr="https://atsdo.xyz"

# echo "File upload test";
# curl ${addr}/upload -F estiemail=boo@gmail.com -F estiphoto=@/home/charliepi/Downloads/debian.jpg;
# echo "\tshould be '0'\n";

# echo "Get Photo by Email";
# curl ${addr}/getphotobyemail?entry=boo@gmail.com;
# echo "\tshould be a b64image";

echo "\nLogin Test";
curl ${addr}/login -F token=atsadmin2 -F email=porthose.cjsmo.cjsmo@gmail.com -F pwd=porthose;
echo "\tshould be true\n";

echo "Login Test Fail";
curl ${addr}/login -F token=atsadmin -F email=porthose.cjsmo.cjsmo@gmail.com -F pwd=porthose;
echo "\tshould be false\n";

# echo "Insert Rev Test";
# curl ${addr}/ins_rev?entry=booSPLITboo@gmail.comSPLITgoodSPACEjogSPLIT5;
# echo "\tshould be ["0", "0"]\n";

# echo "Insert Est Test";
# curl ${addr}/ins_est?entry=booSPACEfuckSPLIT789SPACEhullSPACEaveSPLITportSPACEorchardSPLIT456-456-4566SPLITbooATgmailDOTcomSPLIT07-09-2023SPLITgoodSPACEjob;
# echo "\tshould be ["0", "0"]\n";

# echo "All_Est Test";
# curl ${addr}/all_est;
# echo "\tshould be a populated array\n";

# echo "All_Revs Test";
# curl ${addr}/all_revs;
# echo "\tshould be a populated array\n";

# echo "Revs Backup Test";
# curl ${addr}/revbup;
# echo "\tshould be a "Backup Created"\n";

# echo "Est Backup Test";
# curl ${addr}/estbup;
# echo "\tshould be a "Backup Created"\n";


