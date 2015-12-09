import React from 'react';
// Needed for React Developer Tools
window.React = React;

import {render} from 'react-dom';
import {Router, Route, Link} from 'react-router';

import Master from './Master';
import Song from './song';

render((
  <Router>
    <Route path="/" component={Master}>
      <Route path="/song/:songId" component={Song}/>
    </Route>
  </Router>
),
document.body);
