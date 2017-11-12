function addElement(data) {
    var element = $('<li> <span></span> <button class="delete-button">Удалить</button> </li>');
    $('span', element).text(data);
    $('#root ul').append(element);
    
    $('.delete-button', element).click(
        function(event) {
            $(this).parent().remove()
        }
    );
}

$(function(){
    $('#root').append('<ul></ul>');
    $('#root').append('<input id="add_task_input"> <button id="add_task">Добавить</button>');
    
    $('#add_task').click(
        function() {
            addElement($('#add_task_input').val())
        }
    );
    
    addElement('Сделать задание #3 по web-программированию');
})
