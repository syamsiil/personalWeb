let name = "Seeus";

function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "") {
    return alert("Name Required!");
  } else if (email == "") {
    return alert("Email Required!");
  } else if (phone == "") {
    return alert("Phone Number Required!");
  } else if (subject == "") {
    return alert("Subject Required!");
  } else if (message == "") {
    return alert("Message Required!");
  }

  let emailReceiver = "muhammadalisyamsi@gmail.com";

  //   <a href="mailto:muhammadalisyamsi@gmail.com?subject=frontend&body=Helo, Its me Seeus, What can i do for you?. lets connect with me to 082125073396, Thankyou."></a>

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Helo, Its me ${name}, ${message}. What can i do for you?. Lets connect with me to ${phone}, Thankyou.`;
  a.click();

  console.log(name);
  console.log(email);
  console.log(phone);
  console.log(subject);
  console.log(message);

  let emailer = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(emailer);
}
