import styles from './styles/app.css';
import jQuery from 'jquery';

window.$ = jQuery;
window.jQuery = jQuery;

$(document).ready(function () {
    $('.day-switcher').on('click', function () {
        const currentDayBlock = $(this).closest('.day');

        const direction = $(this).data('switch-direction');
        const targetDay = direction === 'back' ? currentDayBlock.data('previous-day') : currentDayBlock.data('next-day');

        const nextDay = $('.day[data-day="' + targetDay + '"]');
        nextDay.removeClass('hidden');
        currentDayBlock.addClass('hidden');
    });
});