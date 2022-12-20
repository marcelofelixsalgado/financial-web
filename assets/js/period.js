$('#period-form').on('submit', period);

$('#update-period').on('click', updatePeriod);
$('.delete-period').on('click', deletePeriod);

function period(event) {
    event.preventDefault();

    $.ajax({
        url: "/periods",
        method: "POST",
        data: {
            code: $('#code').val(),
            name: $('#name').val(),
            year: $('#year').val(),
            start_date: $('#start_date').val(),
            end_date: $('#end_date').val(),
        }
    }).done(function() {
        window.location = "/periods";
    }).fail(function(error) {
        alert("erro!")
    });    
}

function updatePeriod() {
    $(this).prop('disabled', true);

    const periodId = $(this).data('period-id');
    
    $.ajax({
        url: `/periods/${periodId}`,
        method: "PUT",
        data: {
            code: $('#code').val(),
            name: $('#name').val(),
            year: $('#year').val(),
            start_date: $('#start_date').val(),
            end_date: $('#end_date').val(),
        }
    }).done(function() {
        // Swal.fire('Sucesso!', 'Período criado com sucesso!', 'success')
        //     .then(function() {
        //         window.location = "/home";
        //     })
        alert("Sucesso!");
        window.location = "/periods";
    }).fail(function() {
        // Swal.fire("Ops...", "Erro ao editar o período!", "error");
        alert("Erro ao editar o período!")
    }).always(function() {
        $('#update-period').prop('disabled', false);
    })
}

function deletePeriod(evento) {
    evento.preventDefault();

    // Swal.fire({
    //     title: "Atenção!",
    //     text: "Tem certeza que deseja excluir esse período? Essa ação é irreversível!",
    //     showCancelButton: true,
    //     cancelButtonText: "Cancelar",
    //     icon: "warning"
    // }).then(function(confirmacao) {
    //     if (!confirmacao.value) return;

        const clickedElement = $(evento.target);
        const period = clickedElement.closest('div')
        const periodId = period.data('period-id');
    
        clickedElement.prop('disabled', true);
    
        $.ajax({
            url: `/periods/${periodId}`,
            method: "DELETE"
        }).done(function() {
            period.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            // Swal.fire("Ops...", "Erro ao excluir o período!", "error");
            alert("Erro ao excluir o período!")
        });
    // })

}