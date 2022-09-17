function setModalValues(task) {
    var title = document.getElementById('modal-input-title');
    var desc = document.getElementById('modal-input-desc');
    var priorities = document.querySelectorAll('input[name="task-priority"]');
    var submit = document.getElementById('modal-submit');

    title.value = task ? task.Title || '' : '';
    desc.value =  task ? task.Description || '' : '';
    submit.innerText = task ? "Save changes" : 'Add task';
    
    var ok = false;

    if (task) {
        for (var i = 0; i < priorities.length; i++) {
            if (priorities[i].value === task.Priority) {
                priorities[i].checked = true;
                ok = true;
                break;
            }
        }
    }

    if (!ok || !task.Priority) {
        priorities[0].checked = true;
    }
}

function openModal(task) {
    var title = document.getElementById('modal-input-title');
    var desc = document.getElementById('modal-input-desc');
    var priorities = document.querySelectorAll('input[name="task-priority"]');
    var submit = document.getElementById('modal-submit');

    var original = structuredClone(task);

    submit.onclick = (
        task
        ? () => submitChange(original, task)
        : () => submitChange(task, undefined)
    );

    setModalValues(task);

    if (!task) {
        task = {
            Title: '',
            Description: '',
            Status: 'pending',
            Priority: 'low',
        };
    }

    title.onchange = (event) => task['Title'] = event.target.value;
    desc.onchange = (event) => task['Description'] = event.target.value;

    for (var i = 0; i < priorities.length; i++) {
        priorities[i].onclick = (event) => task['Priority'] = event.target.value;
    }

    var wrapper = document.getElementsByClassName('modal-edit')[0];
    wrapper.style.display = 'block';
}

function closeModal(event) {
    var modal = document.getElementsByClassName('modal-edit')[0];
    var close = document.getElementById('modal-close-button');

    if (event.target == modal || event.target == close) {
        setModalValues(undefined);
        modal.style.display = 'none';
    }
}

function submitChange(old, n) {
    var opts;
    var url = '/api/';
    var message = '';
    
    if (n) {
        opts = {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json'
            },
    
            body: JSON.stringify([old, n]),
        };

        url += 'edit';

        message += 'Changes saved successfully';
        console.log(old, n)
    } else {
        opts = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
    
            body: JSON.stringify(old),
        };

        url += 'add';

        message += 'Task \"' + old.Title + '\" added successfully';
        console.log("add")
    }

    fetch(url, opts)
        .then(res => {
            if (res.status >= 300)
                throw new Error(res.statusText);

            window.alert(message);
            location.reload();
        })
        .catch(err => window.alert('Could not save changes:\n' + err));
        
}

function updateTask(task) {
    var status = 'pending';

    if (task.Status) {
        switch (task.Status) {
            case 'doing':
                status = 'done';
                break;
            case 'done':
                status = 'pending';
                break;
            default:
                status = 'doing';
                break;
        }
    }
    
    task.Status = status;
    fetch('/api/update', {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },

        body: JSON.stringify({ Status: status, ...task }),
    })
        .then(res => {
            if (res.status !== 200)
                throw new Error(res.statusText);

            window.alert(`Task '${task.Title}' updated to '${task.Status}'`);
            location.reload();
        })
        .catch(err => window.alert('Could not update:\n' + err));
}

function generateHTMLForTask(task) {
    var wrapper, title, body, header, badges, buttonWrapper, desc;
    var buttons = [
        {
            title: 'Edit',
            class: 'task-edit',
            el: null,
        },
        {
            title: 'Remove',
            class: 'task-remove',
            el: null,
        }
    ];

    // wrapper
    wrapper = document.createElement('li');
    wrapper.className = 'task';
    wrapper.id = `task-${task.Title.toLowerCase()}`;

    // badges -> status and priority
    badges = document.createElement('ul');
    badges.className = 'task-badges';
    
    temp = document.createElement('li');
    temp.innerText = task.Priority || 'low';
    temp.className = `task-priority-${task.Priority || 'low'}`;

    badges.appendChild(temp);
    
    temp = document.createElement('li');
    temp.innerText = task.Status || 'pending';
    temp.className = 'task-status-' + temp.innerText;

    temp.onclick = () => updateTask(task);

    badges.appendChild(temp);

    // title
    title = document.createElement('h2');
    title.innerText = task.Title || '...';
    title.className = 'task-title';
    
    // header -> title + badges
    header = document.createElement('div');
    header.className = 'task-header';
    header.appendChild(title);
    header.appendChild(badges);
    
    
    body = document.createElement('div');
    body.className = 'task-body';

    buttonWrapper = document.createElement('div');
    buttonWrapper.className = 'task-action-wrapper';
    
    desc = document.createElement('p');
    desc.innerText = task.Description || '';
    desc.classList = 'task-description';
    body.appendChild(desc);

    buttons.forEach((b, index) => {
        buttons[index].el = document.createElement('button');
        buttons[index].el.onclick = (b.title.toLowerCase() === 'edit')
            ? () => openModal(task)
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
    var container = document.getElementById('task-list');
    var add = document.getElementById('add-task');
    add.onclick = () => openModal(undefined);
    
    if (tasks.length && tasks.length > 0) {

        tasks.forEach(task => {
            container.appendChild(generateHTMLForTask(task));
        });
    } else {
        var div = document.createElement('h3');
        div.style = 'text-align: center;';
        div.innerText = 'Empty...';
        
        container.appendChild(div);
    }
}

function removeTask(title) {
    var init = {
        method: 'DELETE',
    };

    fetch('/api/remove?title=' + title, init)
        .then(res => {
            location.reload()
        })
        .catch(err => console.error(err));
}

function getTasks() {
    fetch('/api/get')
        .then(res => res.json())
        .then(json => {
            if (json.content && json.content.length > 0) {
                renderTasks(json.content);
            }
        })
        .catch(err => {
            console.error(err);
            renderTasks([]);
        });
}

window.onclick = closeModal;

getTasks();
