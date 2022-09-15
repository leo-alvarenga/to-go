var host = "127.0.0.1" + window.location.host.split(":")[1]
const init  = {
    method: "GET",
    credentials: "same-origin",
}

var tasks = [
    {
        Id: "1",
        Title: "Okay",
        Description: "ok",
        Priority: "high",
        Status: "pending",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "a",
        Description: "ok",
        Priority: "medium",
        Status: "doing",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "b",
        Description: "ok",
        Priority: "low",
        Status: "done",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "Okay",
        Description: "ok",
        Priority: "high",
        Status: "pending",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "a",
        Description: "ok",
        Priority: "medium",
        Status: "doing",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "b",
        Description: "ok",
        Priority: "low",
        Status: "done",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "Okay",
        Description: "ok",
        Priority: "high",
        Status: "pending",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "a",
        Description: "ok",
        Priority: "medium",
        Status: "doing",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    },
    {
        Id: "1",
        Title: "b",
        Description: "ok",
        Priority: "low",
        Status: "done",
        CreatedIn:"September 12nd, 2022",
        FinishedIn: "",
    }
];

function removeTask(title) {
    var init = {
        method: 'DELETE',
    };

    fetch("/api/remove?title=" + title, init)
        .then(res => {
            console.log(res.status)
            location.reload()
        })
        .catch(err => console.error(err));
}

function generateHTMLForTask(task) {
    var wrapper, title, body, header, badges, buttonWrapper, desc;
    var buttons = [
        {
            title: "Edit",
            class: "task-edit",
            el: null,
        },
        {
            title: "Remove",
            class: "task-remove",
            el: null,
        }
    ];

    // wrapper
    wrapper = document.createElement('li');
    wrapper.className = "task";
    wrapper.id = `task-${task.Title.toLowerCase()}`;

    // badges -> status and priority
    badges = document.createElement('ul');
    badges.className = "task-badges";
    
    temp = document.createElement('li');
    temp.innerText = task.Priority || "low";
    temp.className = `task-priority-${task.Priority || "low"}`;

    badges.appendChild(temp);
    
    temp = document.createElement('li');
    if (task.Status && task.Status === "pending") {
        temp.innerText = task.Status || "pending";
        temp.className = "task-status-pending";
    } else {
        var check = document.createElement('input');

        check.type = "checkbox"
        check.className = "task-check";
        check.checked = task.Status === "done";
        check.onchange = (event) => console.log(event.target.value);

        temp.appendChild(check);
    }

    badges.appendChild(temp);

    // title
    title = document.createElement('h2');
    title.innerText = task.Title || "...";
    title.className = "task-title";
    
    // header -> title + badges
    header = document.createElement('div');
    header.className = "task-header";
    header.appendChild(title);
    header.appendChild(badges);
    
    
    body = document.createElement('div');
    body.className = "task-body";

    buttonWrapper = document.createElement('div');
    buttonWrapper.className = "task-action-wrapper";
    
    desc = document.createElement('p');
    desc.innerText = task.Description || "";
    desc.classList = "task-description";
    body.appendChild(desc);

    buttons.forEach((b, index) => {
        buttons[index].el = document.createElement('button');
        buttons[index].el.onclick = (b.title === "edit")
            ? () => null
            : () => removeTask(task.Title);

        buttons[index].el.innerText = buttons[index].title;
        buttons[index].el.className = buttons[index].class;
        buttonWrapper.appendChild(buttons[index].el);
    });

    body.appendChild(buttonWrapper);
    wrapper.appendChild(header);
    wrapper.appendChild(body);

    return wrapper;
}

function renderTasks(tasks) {
    var container = document.getElementById("task-list");
    
    if (tasks.length && tasks.length > 0) {

        tasks.forEach(task => {
            container.appendChild(generateHTMLForTask(task));
        });
    } else {
        var div = document.createElement('h3');
        div.style = "text-align: center;";
        div.innerText = "Empty...";
        
        container.appendChild(div);
    }
}

function getTasks() {
    fetch("/api/get", init)
        .then(res => res.json())
        .then(json => {
            if (json.content && json.content.length > 0) {
                renderTasks(json.content);
            }
        })
        .catch(err => {
            console.error(err);
            renderTasks(tasks);
        });
}

getTasks();
