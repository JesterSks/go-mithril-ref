/*
   Copyright 2021 Robert Burghart

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
const m = require("mithril")
const i18n = require("i18next").default
const Thing = require("../models/Thing")

module.exports = {
    oninit: Thing.loadThing,
    view: function() {
        let t = Thing.data
        return m("div", [
            m("div.hero.hero-sm.bg-primary",
                m("div.hero-body", [
                    m("h1", "Hero Title"),
                    m("p", "This is a hero example")
                ])
            ),
            m("div.container",
                m("div.columns.col-gapless",
                    m("div.column.col-12",[
                        m("p", t["test"]),
                        m("p", i18n.t("key")),
                    ])
                )
            )
        ])
    }
}
