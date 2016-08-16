var nodemailer = require('nodemailer');
var fs = require('fs');


var img = fs.readFileSync(__dirname+"/cat.jpg");
var attachment = [{
     'filename':'cat.jpg',
     'contents': img
}]

var transport = nodemailer.createTransport("SMTP", {
    host: 'localhost',
});

//var send_email = function (email_content) {
var mailOptions = {
        from: 'nuaays@gmail.com',
        to: 'yangsen@zhongan.com',
        subject: 'Hello, cat!',
        attachments: attachment,
        html:"<b>thanks a for visiting!</b> 世界，你好！" 
};

transport.sendMail(mailOptions, function (error, info) {
        if (error) {
            console.log(error);
        } else {
            console.log('Message sent: ' + info.message);
        }
        transport.close();
});

//https://howtonode.org/sending-e-mails-with-node-and-nodemailer

//https://github.com/nodemailer/nodemailer
//http://www.cnblogs.com/pingfan1990/p/4864822.html
//http://www.nodejs.net/a/20130219/220632.html 
