'use strict';

$('#sidebar-btn').click(function () {
  $('#sidebar').toggleClass('open');
  $('#main').toggleClass('sidebar-open');
});

function fixSidebarHeight() {
  $('#sidebar').height($('body').height());
}

// Fix the sidebar height
fixSidebarHeight();

// Call all functions that need to be called on window resize when 
// the window resizes
$(window).resize(function () {
  fixSidebarHeight();
});
// Constantly check if the window size changed
var $window = $(window),
    wHeight = $window.height();
setInterval(function () {
  if (wHeight != $window.height()) {
    wHeight = $window.height();
    fixSidebarHeight();
  }
}, 300);