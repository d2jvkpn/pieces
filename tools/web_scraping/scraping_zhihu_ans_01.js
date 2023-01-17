var target = document.querySelectorAll(".AnswerCard")[0];

var link = document.createElement("a");
link.href = `data:text/plain;charset=utf8,${document.URL}\n${target.innerText.replace(/\n/g, "%0D%0A")}\n`;

link.download = `${document.title}.md`;
link.click();
