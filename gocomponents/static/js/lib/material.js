'use strict';

$('#sidebar-btn').click(function () {
  $('#sidebar').toggleClass('open');
  $('#main').toggleClass('sidebar-open');
});