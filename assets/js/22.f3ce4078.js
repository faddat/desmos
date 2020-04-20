(window.webpackJsonp=window.webpackJsonp||[]).push([[22],{231:function(e,s,a){"use strict";a.r(s);var t=a(0),r=Object(t.a)({},(function(){var e=this,s=e.$createElement,a=e._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"query-the-stored-posts"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#query-the-stored-posts"}},[e._v("#")]),e._v(" Query the stored posts")]),e._v(" "),a("p",[e._v("This query endpoint allows you to get all the stored posts that match one or more filters.")]),e._v(" "),a("p",[a("strong",[e._v("CLI")])]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("desmoscli query posts "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("[")]),e._v("--flags"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("]")]),e._v("\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("Available flags:")]),e._v(" "),a("ul",[a("li",[a("code",[e._v("--parent-id")]),e._v(" (e.g. "),a("code",[e._v("--parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("--creation-time")]),e._v(" (e.g. "),a("code",[e._v("--creation-time=2020-01-01T12:00:00")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("--allows-comments")]),e._v(" (e.g. "),a("code",[e._v("--allows-comments=true")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("--subspace")]),e._v(" (e.g. "),a("code",[e._v("--subspace=desmos")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("--creator")]),e._v(" (e.g. "),a("code",[e._v("--creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("--sort-by")]),e._v(" (e.g. "),a("code",[e._v("--sort-by=created")]),e._v(")"),a("br"),e._v("\nAccepted values:\n"),a("ul",[a("li",[a("code",[e._v("created")])]),e._v(" "),a("li",[a("code",[e._v("id")]),e._v(" (default)")])])]),e._v(" "),a("li",[a("code",[e._v("--sort-order")]),e._v(" (e.g. "),a("code",[e._v("--sort-order=descending")]),e._v(")"),a("br"),e._v("\nAccepted values:\n"),a("ul",[a("li",[a("code",[e._v("ascending")])]),e._v(" "),a("li",[a("code",[e._v("descending")])])])])]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# Example")]),e._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# desmoscli query posts --parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc --allows-comments=true --subspace=desmos --sort=created --sort-order=descending")]),e._v("\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br")])]),a("p",[a("strong",[e._v("REST")])]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("/posts\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("Available parameters:")]),e._v(" "),a("ul",[a("li",[a("code",[e._v("parent_id")]),e._v(" (e.g. "),a("code",[e._v("parent_id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("creation_time")]),e._v(" (e.g. "),a("code",[e._v("creation_time=2020-01-01T12:00:00")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("allows_comments")]),e._v(" (e.g. "),a("code",[e._v("allows_comments=true")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("subspace")]),e._v(" (e.g. "),a("code",[e._v("subspace=desmos")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("creator")]),e._v(" (e.g. "),a("code",[e._v("creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("sort_by")]),e._v(" (e.g. "),a("code",[e._v("sort_by=created")]),e._v(")")]),e._v(" "),a("li",[a("code",[e._v("sort_order")]),e._v(" (e.g. "),a("code",[e._v("sort_order=descending")]),e._v(")")])]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# Example")]),e._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# curl http://lcd.morpheus.desmos.network:1317/posts?parent_id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc&allows_comments=true&subspace=desmos&sort_by=created&sort_order=descending")]),e._v("\n\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br"),a("span",{staticClass:"line-number"},[e._v("3")]),a("br")])])])}),[],!1,null,null,null);s.default=r.exports}}]);