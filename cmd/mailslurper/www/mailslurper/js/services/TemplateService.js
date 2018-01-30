// Copyright 2013-2016 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

"use strict";

window.TemplateService = {
	load: function (name) {
		return new Promise(function (resolve, reject) {
			$.get("/www/mailslurper/templates/" + name + ".hbs").then(
				function (result) {
					return resolve(result);
				},
				function (xhr, errorType, err) {
					return reject(err);
				}
			);
		});
	}
};

