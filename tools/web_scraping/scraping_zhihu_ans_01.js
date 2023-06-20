var target = document.querySelectorAll(".AnswerCard")[0];

// var target = document.querySelectorAll(".MoreAnswers")[0];

var data = `${document.URL}\n\n${target.innerText}`;

var link = document.createElement("a");
link.href = `data:text/plain;charset=utf8,${data.replace(/\n/g, "%0D%0A")}\n`;

link.download = `${document.title}.md`;
link.click();
