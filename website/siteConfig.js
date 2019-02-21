/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// See https://docusaurus.io/docs/site-config for all the possible
// site configuration options.

// List of projects/orgs using your project for the users page.

const siteConfig = {
  title: 'Practical Common Lisp', // Title for your website.
  tagline: 'A copy of Practical Common Lisp by Peter Seibel',
  url: 'https://practical-common-lisp.plankton.technology', // Your website URL
  baseUrl: '/', // Base URL for your project */
  // For github.io type URLs, you would set the url and baseUrl like:
  //   url: 'https://facebook.github.io',
  //   baseUrl: '/test-site/',

  // Used for publishing and more
  projectName: 'plankton.technology',
  organizationName: 'necroplankton',
  // For top-level user or org sites, the organization is still the same.
  // e.g., for the https://JoelMarcey.github.io site, it would be set like...
  //   organizationName: 'JoelMarcey'

  // For no header links in the top nav bar -> headerLinks: [],
  headerLinks: [
    {doc: 'introduction-why-lisp', label: 'Docs'}
  ],

  // If you have users set above, you add it here:

  /* path to images for header/footer */
  headerIcon: 'img/common-lisp.svg',
  footerIcon: 'img/common-lisp.svg',
  favicon: 'img/common-lisp.svg',

  /* Colors for website */
  colors: {
    primaryColor: '#00bcd4',
    secondaryColor: '#607d8b',
  },

  /* Custom fonts for website */
  /*
  fonts: {
    myFont: [
      "Times New Roman",
      "Serif"
    ],
    myOtherFont: [
      "-apple-system",
      "system-ui"
    ]
  },
  */

  // This copyright info is used in /core/Footer.js and blog RSS/Atom feeds.
  copyright: `Copyright © ${new Date().getFullYear()} Peter Seibel`,

  highlight: {
    // Highlight.js theme to use for syntax highlighting in code blocks.
    theme: 'atelier-estuary-dark',
  },

  // Add custom scripts here that would be placed in <script> tags.
  scripts: ['https://buttons.github.io/buttons.js'],

  // On page navigation for the current documentation page.
  onPageNav: 'separate',
  // No .html extensions for paths.
  cleanUrl: true,

  // Open Graph and Twitter card images.
  ogImage: 'img/book.gif',
  twitterImage: 'img/book.gif',

  // Show documentation's last contributor's name.
  // enableUpdateBy: true,

  // Show documentation's last update time.
  // enableUpdateTime: true,

  // You may provide arbitrary config keys to be used as needed by your
  // template. For example, if you need your repo's URL...
  //   repoUrl: 'https://github.com/facebook/test-site',
};

module.exports = siteConfig;
