# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0](https://github.com/iB2/whatsapp-mcp/compare/v0.6.0...v0.1.0) (2026-09-02)


### release

* adopt Release Please for automated versioning/changelog ([#15](https://github.com/iB2/whatsapp-mcp/issues/15)) ([bef1a96](https://github.com/iB2/whatsapp-mcp/commit/bef1a966ef254947f37833d06549665901133890))
* adopt Release Please for automated versioning/changelog ([#15](https://github.com/iB2/whatsapp-mcp/issues/15)) ([6d45958](https://github.com/iB2/whatsapp-mcp/commit/6d45958139effa3079ff27a9708d400f89ba9ddf))


### Features

* add chat filter, cleanup script, and analyze_chats tool ([8e41125](https://github.com/iB2/whatsapp-mcp/commit/8e411253d193e7c64ee585d2b3abcf965ae8e5ac))
* add explicit message read receipts ([#201](https://github.com/iB2/whatsapp-mcp/issues/201)) ([e35224e](https://github.com/iB2/whatsapp-mcp/commit/e35224e182fe06632ae0436949dda3cded544b17))
* add image media support in webhook forwarding ([#45](https://github.com/iB2/whatsapp-mcp/issues/45)) ([43d7794](https://github.com/iB2/whatsapp-mcp/commit/43d7794d8e003dc72fb9dba3b55eacee479753c8))
* add image media support in webhook forwarding ([#45](https://github.com/iB2/whatsapp-mcp/issues/45)) ([c5f409c](https://github.com/iB2/whatsapp-mcp/commit/c5f409c78da8b7029b418e2fa428316dad7f970b))
* **bridge:** --full-history-pair flag to request full history at pair time ([#37](https://github.com/iB2/whatsapp-mcp/issues/37)) ([59b834f](https://github.com/iB2/whatsapp-mcp/commit/59b834fcba2b0d9994d6a21946aae59ab65cff99))
* **bridge:** --full-history-pair flag to request full history at pair time ([#37](https://github.com/iB2/whatsapp-mcp/issues/37)) ([2358dd5](https://github.com/iB2/whatsapp-mcp/commit/2358dd5cf8cda17110076446c73637eab23bc7c3))
* **bridge:** add macOS launchd installer ([#116](https://github.com/iB2/whatsapp-mcp/issues/116)) ([dbc63ae](https://github.com/iB2/whatsapp-mcp/commit/dbc63ae81b6c0d1e11fa9a06122522f1ce00e56b))
* **bridge:** add on-demand history sync for a single chat ([#168](https://github.com/iB2/whatsapp-mcp/issues/168)) ([f44440b](https://github.com/iB2/whatsapp-mcp/commit/f44440b5874dba67c3daff2709458b2c45d7e1f7))
* **bridge:** add outbound webhook opt-out ([#204](https://github.com/iB2/whatsapp-mcp/issues/204)) ([6c3f3fe](https://github.com/iB2/whatsapp-mcp/commit/6c3f3fe2763584879166b5f92834cf7635d16a58))
* **bridge:** add support for stickers ([#110](https://github.com/iB2/whatsapp-mcp/issues/110)) ([5d98686](https://github.com/iB2/whatsapp-mcp/commit/5d98686b6007c768cefcbd314f2b2232724d0db4))
* **bridge:** capture incoming WhatsApp call events ([#39](https://github.com/iB2/whatsapp-mcp/issues/39)) ([4f1664a](https://github.com/iB2/whatsapp-mcp/commit/4f1664a9e54858e0b48f1d417fd4aee9ea3bbf35))
* **bridge:** capture incoming WhatsApp call events ([#39](https://github.com/iB2/whatsapp-mcp/issues/39)) ([197a0c9](https://github.com/iB2/whatsapp-mcp/commit/197a0c9de304b762db3e3a7fb98d9dd4ce510f74))
* **bridge:** configurable linked-device name via WHATSAPP_DEVICE_NAME ([#157](https://github.com/iB2/whatsapp-mcp/issues/157)) ([1a71032](https://github.com/iB2/whatsapp-mcp/commit/1a71032fd878c19eea346292f147350fadb4d818)), closes [#156](https://github.com/iB2/whatsapp-mcp/issues/156)
* **bridge:** forward reaction webhook events ([#129](https://github.com/iB2/whatsapp-mcp/issues/129)) ([2c8062f](https://github.com/iB2/whatsapp-mcp/commit/2c8062fc1990580fe2172b0d88bc1db30d0100cb))
* **bridge:** native polls with encrypted vote decoding ([#1](https://github.com/iB2/whatsapp-mcp/issues/1)) ([2eb0867](https://github.com/iB2/whatsapp-mcp/commit/2eb08670d38bbee44242e1be7c0ca26e1289b30e))
* **bridge:** persist chat read state from read receipts + history-sync backfill ([#155](https://github.com/iB2/whatsapp-mcp/issues/155)) ([6c6ec00](https://github.com/iB2/whatsapp-mcp/commit/6c6ec006a1e67d236172425dbb96a74d6cb9c7c3))
* capture inbound reactions and add /api/react endpoint ([#108](https://github.com/iB2/whatsapp-mcp/issues/108)) ([58645d5](https://github.com/iB2/whatsapp-mcp/commit/58645d53c57d619dd22454e4b52ab0c41538f445)), closes [#106](https://github.com/iB2/whatsapp-mcp/issues/106)
* **mcp:** add @-mention support to send_message ([#190](https://github.com/iB2/whatsapp-mcp/issues/190)) ([b86a57d](https://github.com/iB2/whatsapp-mcp/commit/b86a57d3711f96ba3750d6a4eabf96f8279473b0))
* **mcp:** support http and sse transports via env var ([#112](https://github.com/iB2/whatsapp-mcp/issues/112)) ([f885c6c](https://github.com/iB2/whatsapp-mcp/commit/f885c6c37a8512f6df7437c0779b32e10fb1823c))
* optional caption on send_file ([#193](https://github.com/iB2/whatsapp-mcp/issues/193)) ([9f6324c](https://github.com/iB2/whatsapp-mcp/commit/9f6324c92f6a30e3231d8337871e206014a2b66a))
* persist inbound quoted_message_id and add reply support to /api/send ([#109](https://github.com/iB2/whatsapp-mcp/issues/109)) ([731e3eb](https://github.com/iB2/whatsapp-mcp/commit/731e3eb46ab94200396d92607c6b8e8f27c306d9)), closes [#107](https://github.com/iB2/whatsapp-mcp/issues/107)


### Bug Fixes

* **bridge:** authenticate outbound webhook POSTs ([3e8b919](https://github.com/iB2/whatsapp-mcp/commit/3e8b919290310fbaedfe9bcfa29f03f8336cd1e2))
* **bridge:** auto-download runs after StoreMessage to avoid lookup race ([#41](https://github.com/iB2/whatsapp-mcp/issues/41)) ([6a711b7](https://github.com/iB2/whatsapp-mcp/commit/6a711b72740b44715c0c7e94ac4c432ee533ff32))
* **bridge:** auto-download runs after StoreMessage to avoid lookup race ([#41](https://github.com/iB2/whatsapp-mcp/issues/41)) ([0c267f5](https://github.com/iB2/whatsapp-mcp/commit/0c267f5f0834138f54d7e2d731ffbd8f8deb5929))
* **bridge:** bump whatsmeow for client compatibility ([#182](https://github.com/iB2/whatsapp-mcp/issues/182)) ([544e9e7](https://github.com/iB2/whatsapp-mcp/commit/544e9e737e980cf652794262fb0f3957e5c10fa9))
* **bridge:** bump whatsmeow so WhatsApp accepts new device pairing ([#128](https://github.com/iB2/whatsapp-mcp/issues/128)) ([c5cc5c7](https://github.com/iB2/whatsapp-mcp/commit/c5cc5c78dcdc64ded482bf451507a785f5afbed6))
* **bridge:** document filename leaked the sender's absolute path ([#2](https://github.com/iB2/whatsapp-mcp/issues/2)) ([3b3d4e5](https://github.com/iB2/whatsapp-mcp/commit/3b3d4e57b4fa1e962a1df7c097b8460a14d73727))
* **bridge:** extract text from Template, Button, Interactive, and List messages ([#134](https://github.com/iB2/whatsapp-mcp/issues/134)) ([a76645a](https://github.com/iB2/whatsapp-mcp/commit/a76645ae8ca5a80e3e388fd62751c45a3a23933d))
* **bridge:** fall back to local whatsmeow_contacts for chat name resolution ([#135](https://github.com/iB2/whatsapp-mcp/issues/135)) ([839d069](https://github.com/iB2/whatsapp-mcp/commit/839d069ec95902a433882cf75cc5d26e3710686e))
* **bridge:** forward native WhatsApp activation metadata ([#173](https://github.com/iB2/whatsapp-mcp/issues/173)) ([b823a9e](https://github.com/iB2/whatsapp-mcp/commit/b823a9e282356c734a9ae3d36291a76a46102603))
* **bridge:** handle ProtocolMessage_REVOKE (delete-for-everyone) events ([#99](https://github.com/iB2/whatsapp-mcp/issues/99)) ([ab08698](https://github.com/iB2/whatsapp-mcp/commit/ab08698eb56f19d6b360b5c452fcab280d3d33cd))
* **bridge:** handle ProtocolMessage_REVOKE (delete-for-everyone) events ([#99](https://github.com/iB2/whatsapp-mcp/issues/99)) ([7f4ec42](https://github.com/iB2/whatsapp-mcp/commit/7f4ec42c455d1716e901b6262a9ff7e3af437b6a))
* **bridge:** handle StreamReplaced event to recover from session conflicts ([#27](https://github.com/iB2/whatsapp-mcp/issues/27)) ([27a4b5c](https://github.com/iB2/whatsapp-mcp/commit/27a4b5c42a85438ab806cf27ab5032f830fdb433))
* **bridge:** handle StreamReplaced event to recover from session conflicts ([#27](https://github.com/iB2/whatsapp-mcp/issues/27)) ([0cd6475](https://github.com/iB2/whatsapp-mcp/commit/0cd647560705acbbb3d46b236e1e09df7b20ceee))
* **bridge:** include message ID in media filenames to prevent same-second collisions ([#40](https://github.com/iB2/whatsapp-mcp/issues/40)) ([02ce549](https://github.com/iB2/whatsapp-mcp/commit/02ce5494f72e0ada65d6e21e2f7990aa6170d0c5))
* **bridge:** include message ID in media filenames to prevent same-second collisions ([#40](https://github.com/iB2/whatsapp-mcp/issues/40)) ([1e819c1](https://github.com/iB2/whatsapp-mcp/commit/1e819c1d67d7ed087ab3298253f4f987cd3e48e9))
* **bridge:** keep CDN auth tokens in directPath to fix 403 media downloads ([#132](https://github.com/iB2/whatsapp-mcp/issues/132)) ([4e354af](https://github.com/iB2/whatsapp-mcp/commit/4e354afde7321dbd4875ed1a870539e05ef1443d))
* **bridge:** log send caller identity ([#96](https://github.com/iB2/whatsapp-mcp/issues/96)) ([1979b73](https://github.com/iB2/whatsapp-mcp/commit/1979b732d4c6a8ed77e44a91d36a812ff3b4c311))
* **bridge:** log send caller identity ([#96](https://github.com/iB2/whatsapp-mcp/issues/96)) ([c81947c](https://github.com/iB2/whatsapp-mcp/commit/c81947cbd4235e827637961b7ec53b47725d5a65))
* **bridge:** pass message store when sending messages ([#91](https://github.com/iB2/whatsapp-mcp/issues/91)) ([c86b1b2](https://github.com/iB2/whatsapp-mcp/commit/c86b1b210bb1902a683a9e0c443dbe01316c29b7))
* **bridge:** pass message store when sending messages ([#91](https://github.com/iB2/whatsapp-mcp/issues/91)) ([423ada9](https://github.com/iB2/whatsapp-mcp/commit/423ada9cd3432fcde79bb388d527a62cd5e8f696))
* **bridge:** preserve original timestamp on retry-redelivered messages ([#149](https://github.com/iB2/whatsapp-mcp/issues/149)) ([45e674e](https://github.com/iB2/whatsapp-mcp/commit/45e674ef0667303a690bc82de0d8bbc61f4dfd40))
* **bridge:** resolve [@lid](https://github.com/lid) sender to phone JID in webhook payload ([#56](https://github.com/iB2/whatsapp-mcp/issues/56)) ([0a36db4](https://github.com/iB2/whatsapp-mcp/commit/0a36db4c70a2ed423ffe75e764e0d33092b67717))
* **bridge:** resolve [@lid](https://github.com/lid) sender to phone JID in webhook payload ([#56](https://github.com/iB2/whatsapp-mcp/issues/56)) ([746fca0](https://github.com/iB2/whatsapp-mcp/commit/746fca03297438f12b53ad3df4d33a6200773a05))
* **bridge:** set FileName and detect MIME for document sends ([#95](https://github.com/iB2/whatsapp-mcp/issues/95)) ([28aa25c](https://github.com/iB2/whatsapp-mcp/commit/28aa25c77d3ba5e97dc65c47216f57b39e926414))
* **bridge:** set FileName and detect MIME for document sends ([#95](https://github.com/iB2/whatsapp-mcp/issues/95)) ([af18908](https://github.com/iB2/whatsapp-mcp/commit/af18908225767213bfedf970feb68f9a9bd79182))
* **bridge:** surface image/video/document captions in extractTextContent ([#42](https://github.com/iB2/whatsapp-mcp/issues/42)) ([33ef0c0](https://github.com/iB2/whatsapp-mcp/commit/33ef0c057ad47b364f50eba35b074d04a0bd64e1))
* **bridge:** surface image/video/document captions in extractTextContent ([#42](https://github.com/iB2/whatsapp-mcp/issues/42)) ([fbb3f28](https://github.com/iB2/whatsapp-mcp/commit/fbb3f283f0296a5f7e4aaf72eec2b99853de41b8))
* **ci:** make dependabot auto-approve non-fatal ([#65](https://github.com/iB2/whatsapp-mcp/issues/65)) ([b671e82](https://github.com/iB2/whatsapp-mcp/commit/b671e827d2f728f24b5037b525c27694dc6b428f))
* **ci:** make dependabot auto-approve non-fatal ([#65](https://github.com/iB2/whatsapp-mcp/issues/65)) ([ea0dfb1](https://github.com/iB2/whatsapp-mcp/commit/ea0dfb179ac0f68815e5a4b6b96058342590c8bd))
* **deps:** preserve Intel macOS cryptography installs ([#188](https://github.com/iB2/whatsapp-mcp/issues/188)) ([cc43c7b](https://github.com/iB2/whatsapp-mcp/commit/cc43c7b02d0c385cbebedac84dbb4b40103e1fa2))
* escape monitor-script shell vars in launchd installer ([#130](https://github.com/iB2/whatsapp-mcp/issues/130)) ([a3fcaf9](https://github.com/iB2/whatsapp-mcp/commit/a3fcaf94fa31a5824a4b16677586fc71a6556935))
* exit orphaned stdio MCP servers on parent death ([#177](https://github.com/iB2/whatsapp-mcp/issues/177)) ([a6869f4](https://github.com/iB2/whatsapp-mcp/commit/a6869f4fb4c2549efa7317872f1dfa7a553dc67a))
* harden read-state observe/act after [#155](https://github.com/iB2/whatsapp-mcp/issues/155) and [#201](https://github.com/iB2/whatsapp-mcp/issues/201) ([#203](https://github.com/iB2/whatsapp-mcp/issues/203)) ([c3fbd45](https://github.com/iB2/whatsapp-mcp/commit/c3fbd45d6aec4f3a7668f814634910e4ddd3fcd4))
* **mcp:** deduplicate contact chat results ([0da7399](https://github.com/iB2/whatsapp-mcp/commit/0da7399b5ad04b1972a16ceae7abd5fcab4a63a3))
* **mcp:** match messages by both phone number and LID via whatsmeow_lid_map ([#43](https://github.com/iB2/whatsapp-mcp/issues/43)) ([7c4f129](https://github.com/iB2/whatsapp-mcp/commit/7c4f1296e52b065ea6917d2800c081fe3f4debe4))
* **mcp:** match messages by both phone number and LID via whatsmeow_lid_map ([#43](https://github.com/iB2/whatsapp-mcp/issues/43)) ([04c8755](https://github.com/iB2/whatsapp-mcp/commit/04c875568424a125b3e824ba094427f8d899d7c3))
* **mcp:** resolve bare numeric LIDs ([#97](https://github.com/iB2/whatsapp-mcp/issues/97)) ([a8779ab](https://github.com/iB2/whatsapp-mcp/commit/a8779abdccece147cef832166c684086043e808a))
* **mcp:** resolve bare numeric LIDs ([#97](https://github.com/iB2/whatsapp-mcp/issues/97)) ([768dafd](https://github.com/iB2/whatsapp-mcp/commit/768dafd608bc19a4960691917274e90abfd57a02))
* **mcp:** resolve bridge token from whatsmeow db path ([5926173](https://github.com/iB2/whatsapp-mcp/commit/592617391fdfc67f1a44a5471cfed5bcad2cb43f)), closes [#142](https://github.com/iB2/whatsapp-mcp/issues/142)
* **mcp:** resolve contacts via whatsmeow store with LID → phone fallback ([#30](https://github.com/iB2/whatsapp-mcp/issues/30)) ([0c846a4](https://github.com/iB2/whatsapp-mcp/commit/0c846a40ecc8ef3408693a9c0180491e173ba10d))
* **mcp:** resolve contacts via whatsmeow store with LID → phone fallback ([#30](https://github.com/iB2/whatsapp-mcp/issues/30)) ([b9b0175](https://github.com/iB2/whatsapp-mcp/commit/b9b0175e6475d5402feedc190f44758045985992))
* **mcp:** serialize message context result ([af0aed6](https://github.com/iB2/whatsapp-mcp/commit/af0aed6d3a8817ec6f20de34722b5dc482a113f1))
* pin anyio&lt;4.9 to avoid cancel scope regression ([#44](https://github.com/iB2/whatsapp-mcp/issues/44)) ([21a85b7](https://github.com/iB2/whatsapp-mcp/commit/21a85b776218b1872510f65aecc1766152ccfef6))
* pin anyio&lt;4.9 to avoid cancel scope regression ([#44](https://github.com/iB2/whatsapp-mcp/issues/44)) ([627db67](https://github.com/iB2/whatsapp-mcp/commit/627db6746076108ba4ca370ff3f95420ccbb30ef))
* security hardening for LAN exposure and Unicode search ([#55](https://github.com/iB2/whatsapp-mcp/issues/55)) ([05e639b](https://github.com/iB2/whatsapp-mcp/commit/05e639ba8560c0a1475a45f9ae5ecd928058e4fa))
* security hardening for LAN exposure and Unicode search ([#55](https://github.com/iB2/whatsapp-mcp/issues/55)) ([8097f39](https://github.com/iB2/whatsapp-mcp/commit/8097f39dd4a19edd2ebab52d086704b528019005))
* send and track disappearing-message settings correctly ([#82](https://github.com/iB2/whatsapp-mcp/issues/82)) ([e07bd23](https://github.com/iB2/whatsapp-mcp/commit/e07bd2320d8466d02d42cafdf5aaed859948ff8e))
* send and track disappearing-message settings correctly ([#82](https://github.com/iB2/whatsapp-mcp/issues/82)) ([b250409](https://github.com/iB2/whatsapp-mcp/commit/b25040985b0e72d87d091d83d0bb64ff7d78eba8))
* **server:** list_chats/get_chat error when include_last_message=False ([#79](https://github.com/iB2/whatsapp-mcp/issues/79)) ([6cfdb2b](https://github.com/iB2/whatsapp-mcp/commit/6cfdb2bda891dd7de3f359effded7b575fbaaf79))
* **server:** list_chats/get_chat error when include_last_message=False ([#79](https://github.com/iB2/whatsapp-mcp/issues/79)) ([31e03e6](https://github.com/iB2/whatsapp-mcp/commit/31e03e63ddaa5ed2b0037faa7e70cb736addf617))


### Performance Improvements

* **bridge:** index messages(chat_jid) to speed up LID migration ([#144](https://github.com/iB2/whatsapp-mcp/issues/144)) ([609ddc2](https://github.com/iB2/whatsapp-mcp/commit/609ddc2056406dc40d226c4fd4ecded92693dc08))


### Documentation

* add contributors section and updating instructions to README ([#67](https://github.com/iB2/whatsapp-mcp/issues/67)) ([f6fa977](https://github.com/iB2/whatsapp-mcp/commit/f6fa9775992cdb7c347a101baa8d3eef4d684a7b))
* add contributors section and updating instructions to README ([#67](https://github.com/iB2/whatsapp-mcp/issues/67)) ([71b0c03](https://github.com/iB2/whatsapp-mcp/commit/71b0c03070bb2054f45f202590c84d943cb2e4e6))
* add ROADMAP, AGENTS, CONTRIBUTING, CODEOWNERS, issue/PR templates ([#47](https://github.com/iB2/whatsapp-mcp/issues/47)) ([985f8d8](https://github.com/iB2/whatsapp-mcp/commit/985f8d8d22dcbfebda900768aaf8ed9c79d4e6f7))
* add ROADMAP, AGENTS, CONTRIBUTING, CODEOWNERS, issue/PR templates ([#47](https://github.com/iB2/whatsapp-mcp/issues/47)) ([c2e4f5b](https://github.com/iB2/whatsapp-mcp/commit/c2e4f5b22cc3de42b4850222d3cc3ebfa4efbd14))
* document WHATSMEOW_DB_PATH env var ([#51](https://github.com/iB2/whatsapp-mcp/issues/51)) ([558c392](https://github.com/iB2/whatsapp-mcp/commit/558c392ca35126cf3b9e73424e1014c603758bd5))
* document WHATSMEOW_DB_PATH env var ([#51](https://github.com/iB2/whatsapp-mcp/issues/51)) ([c24d151](https://github.com/iB2/whatsapp-mcp/commit/c24d151c9ebe07c052fa1e3a3662c10795747280))
* make AGENTS.md canonical and fix FORWARD_SELF default ([#127](https://github.com/iB2/whatsapp-mcp/issues/127)) ([d31fa06](https://github.com/iB2/whatsapp-mcp/commit/d31fa060cf9fb4a65b8ce0f25744492c1659d72b))
* **readme:** add demo video overview ([#122](https://github.com/iB2/whatsapp-mcp/issues/122)) ([645a7e5](https://github.com/iB2/whatsapp-mcp/commit/645a7e5a46c7511f030cd5ef20f658eb8faca844))
* **readme:** document app state recovery ([#117](https://github.com/iB2/whatsapp-mcp/issues/117)) ([b2cf46b](https://github.com/iB2/whatsapp-mcp/commit/b2cf46bc3bc04b46df7886c8af0bd824b6e3fe9e))
* update remaining "go run main.go" references to "go run ." ([#50](https://github.com/iB2/whatsapp-mcp/issues/50)) ([de75e53](https://github.com/iB2/whatsapp-mcp/commit/de75e53467ba657fb166fda64ad23d8b5cbb9801))
* update remaining "go run main.go" references to "go run ." ([#50](https://github.com/iB2/whatsapp-mcp/issues/50)) ([53e1ca7](https://github.com/iB2/whatsapp-mcp/commit/53e1ca7dd2a470fd5f947f22ee667753cf7e6be6))
* Update SECURITY.md with enhanced security policy ([#78](https://github.com/iB2/whatsapp-mcp/issues/78)) ([8cfaee2](https://github.com/iB2/whatsapp-mcp/commit/8cfaee22c4b4371e0634db50a898017a43fabddd))
* Update SECURITY.md with enhanced security policy ([#78](https://github.com/iB2/whatsapp-mcp/issues/78)) ([214780f](https://github.com/iB2/whatsapp-mcp/commit/214780f3d1cd4c45f04676b37913c1d60b46f445))

## [0.6.0](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.5.1...v0.6.0) (2026-08-11)


### Features

* add explicit message read receipts ([#201](https://github.com/verygoodplugins/whatsapp-mcp/issues/201)) ([65751f1](https://github.com/verygoodplugins/whatsapp-mcp/commit/65751f1ceb32b5db1f421c9b6f73f244de6be2fc))
* **bridge:** add outbound webhook opt-out ([#204](https://github.com/verygoodplugins/whatsapp-mcp/issues/204)) ([a9be700](https://github.com/verygoodplugins/whatsapp-mcp/commit/a9be7001db482c89193c73f0eff620d16883f1c3))
* **bridge:** persist chat read state from read receipts + history-sync backfill ([#155](https://github.com/verygoodplugins/whatsapp-mcp/issues/155)) ([a677828](https://github.com/verygoodplugins/whatsapp-mcp/commit/a677828a180b8ad0bccec83ae81b6471921eac87))
* optional caption on send_file ([#193](https://github.com/verygoodplugins/whatsapp-mcp/issues/193)) ([21df01f](https://github.com/verygoodplugins/whatsapp-mcp/commit/21df01f45afd36b379e00d250974ffb82a1542a3))


### Bug Fixes

* harden read-state observe/act after [#155](https://github.com/verygoodplugins/whatsapp-mcp/issues/155) and [#201](https://github.com/verygoodplugins/whatsapp-mcp/issues/201) ([#203](https://github.com/verygoodplugins/whatsapp-mcp/issues/203)) ([f0892be](https://github.com/verygoodplugins/whatsapp-mcp/commit/f0892beb1dcc23273f85604d287aa741c3fd92a7))

## [0.5.1](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.5.0...v0.5.1) (2026-08-08)


### Bug Fixes

* **deps:** preserve Intel macOS cryptography installs ([#188](https://github.com/verygoodplugins/whatsapp-mcp/issues/188)) ([3ed0538](https://github.com/verygoodplugins/whatsapp-mcp/commit/3ed0538b84ae7ba79613d26a22c6b7be702f647f))

## [0.5.0](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.4.2...v0.5.0) (2026-08-08)


### Features

* **bridge:** add on-demand history sync for a single chat ([#168](https://github.com/verygoodplugins/whatsapp-mcp/issues/168)) ([43afd2e](https://github.com/verygoodplugins/whatsapp-mcp/commit/43afd2ef77aa161ae353cc0d4c318bfd4f071808))
* **bridge:** configurable linked-device name via WHATSAPP_DEVICE_NAME ([#157](https://github.com/verygoodplugins/whatsapp-mcp/issues/157)) ([eb0565d](https://github.com/verygoodplugins/whatsapp-mcp/commit/eb0565de3e123e23e9fcc28a308d34898f0c7d77)), closes [#156](https://github.com/verygoodplugins/whatsapp-mcp/issues/156)
* **mcp:** add @-mention support to send_message ([#190](https://github.com/verygoodplugins/whatsapp-mcp/issues/190)) ([30d2bdb](https://github.com/verygoodplugins/whatsapp-mcp/commit/30d2bdb6947fd87ca21efbe43ae7bfd3c4781ea6))

## [0.4.2](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.4.1...v0.4.2) (2026-08-05)


### Bug Fixes

* **bridge:** authenticate outbound webhook POSTs ([e5f1a9a](https://github.com/verygoodplugins/whatsapp-mcp/commit/e5f1a9aef5c78198ad27d52d40d4513d3b7e0e2f))
* **bridge:** bump whatsmeow for client compatibility ([#182](https://github.com/verygoodplugins/whatsapp-mcp/issues/182)) ([5470342](https://github.com/verygoodplugins/whatsapp-mcp/commit/54703427c5ab25510661cd1bef6f91c49c04f4c5))
* **bridge:** forward native WhatsApp activation metadata ([#173](https://github.com/verygoodplugins/whatsapp-mcp/issues/173)) ([afd2c7a](https://github.com/verygoodplugins/whatsapp-mcp/commit/afd2c7a0a0eba52d7837d02fb7293c662918dde6))
* **bridge:** preserve original timestamp on retry-redelivered messages ([#149](https://github.com/verygoodplugins/whatsapp-mcp/issues/149)) ([bb55a54](https://github.com/verygoodplugins/whatsapp-mcp/commit/bb55a54e36619f4df3cab3fd433cb151b7483d0d))
* exit orphaned stdio MCP servers on parent death ([#177](https://github.com/verygoodplugins/whatsapp-mcp/issues/177)) ([8cc0ecd](https://github.com/verygoodplugins/whatsapp-mcp/commit/8cc0ecd14a949f2c48482086e25a233955390e82))

## [0.4.1](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.4.0...v0.4.1) (2026-06-26)


### Bug Fixes

* **mcp:** resolve bridge token from whatsmeow db path ([000b8e7](https://github.com/verygoodplugins/whatsapp-mcp/commit/000b8e73bea952455180c1914f5e77582568b40a)), closes [#142](https://github.com/verygoodplugins/whatsapp-mcp/issues/142)

## [0.4.0](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.3.0...v0.4.0) (2026-06-26)


### Features

* **bridge:** add support for stickers ([#110](https://github.com/verygoodplugins/whatsapp-mcp/issues/110)) ([f7f475c](https://github.com/verygoodplugins/whatsapp-mcp/commit/f7f475c8ace0d991bea82bea3c75fa4772ae193c))
* **bridge:** forward reaction webhook events ([#129](https://github.com/verygoodplugins/whatsapp-mcp/issues/129)) ([22a03c2](https://github.com/verygoodplugins/whatsapp-mcp/commit/22a03c271ce4f1db19a4c9ff0a46a5770ce7cdff))
* capture inbound reactions and add /api/react endpoint ([#108](https://github.com/verygoodplugins/whatsapp-mcp/issues/108)) ([a7dc663](https://github.com/verygoodplugins/whatsapp-mcp/commit/a7dc663f1ef9c0bb8044304abaf6f1dbee857059)), closes [#106](https://github.com/verygoodplugins/whatsapp-mcp/issues/106)
* **mcp:** support http and sse transports via env var ([#112](https://github.com/verygoodplugins/whatsapp-mcp/issues/112)) ([0a9ba2b](https://github.com/verygoodplugins/whatsapp-mcp/commit/0a9ba2b1f1325b805a6c6b3dac6f0e384aae04ae))


### Bug Fixes

* **bridge:** bump whatsmeow so WhatsApp accepts new device pairing ([#128](https://github.com/verygoodplugins/whatsapp-mcp/issues/128)) ([715be9d](https://github.com/verygoodplugins/whatsapp-mcp/commit/715be9dcfea2edd6d65b59565aa88123a7876353))
* **bridge:** extract text from Template, Button, Interactive, and List messages ([#134](https://github.com/verygoodplugins/whatsapp-mcp/issues/134)) ([d5db52e](https://github.com/verygoodplugins/whatsapp-mcp/commit/d5db52e49633e8e3f08bb24c79d00dcebf0ee535))
* **bridge:** fall back to local whatsmeow_contacts for chat name resolution ([#135](https://github.com/verygoodplugins/whatsapp-mcp/issues/135)) ([37df5b3](https://github.com/verygoodplugins/whatsapp-mcp/commit/37df5b3d895d0b66104543e1649f26ae0b00f2b8))
* **bridge:** keep CDN auth tokens in directPath to fix 403 media downloads ([#132](https://github.com/verygoodplugins/whatsapp-mcp/issues/132)) ([1b67982](https://github.com/verygoodplugins/whatsapp-mcp/commit/1b67982c3a4e51dfbb802802a859739a76039706))
* escape monitor-script shell vars in launchd installer ([#130](https://github.com/verygoodplugins/whatsapp-mcp/issues/130)) ([c758f5f](https://github.com/verygoodplugins/whatsapp-mcp/commit/c758f5f3d88a591e116a6af49f41f07965a9c450))


### Performance Improvements

* **bridge:** index messages(chat_jid) to speed up LID migration ([#144](https://github.com/verygoodplugins/whatsapp-mcp/issues/144)) ([9c6fa66](https://github.com/verygoodplugins/whatsapp-mcp/commit/9c6fa6679e990dcceb12fc80d95bfc3cfce4cca1))


### Documentation

* make AGENTS.md canonical and fix FORWARD_SELF default ([#127](https://github.com/verygoodplugins/whatsapp-mcp/issues/127)) ([8c92254](https://github.com/verygoodplugins/whatsapp-mcp/commit/8c9225450ffd06f6db44445fa3007a4796d04839))
* **readme:** add demo video overview ([#122](https://github.com/verygoodplugins/whatsapp-mcp/issues/122)) ([11672a3](https://github.com/verygoodplugins/whatsapp-mcp/commit/11672a3407f57060cf5efa720a90c9d165fee462))

## [0.3.0](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.2.1...v0.3.0) (2026-06-03)


### Features

* **bridge:** add macOS launchd installer ([#116](https://github.com/verygoodplugins/whatsapp-mcp/issues/116)) ([9b2cb8a](https://github.com/verygoodplugins/whatsapp-mcp/commit/9b2cb8ad3e9e1742fc6ac4714ceb5c4a72fb74e3))
* persist inbound quoted_message_id and add reply support to /api/send ([#109](https://github.com/verygoodplugins/whatsapp-mcp/issues/109)) ([f72d352](https://github.com/verygoodplugins/whatsapp-mcp/commit/f72d35220153dcd8b1fc2dfc645447cf0796271c)), closes [#107](https://github.com/verygoodplugins/whatsapp-mcp/issues/107)


### Bug Fixes

* **bridge:** handle ProtocolMessage_REVOKE (delete-for-everyone) events ([#99](https://github.com/verygoodplugins/whatsapp-mcp/issues/99)) ([7f4ec42](https://github.com/verygoodplugins/whatsapp-mcp/commit/7f4ec42c455d1716e901b6262a9ff7e3af437b6a))
* **mcp:** deduplicate contact chat results ([da543ce](https://github.com/verygoodplugins/whatsapp-mcp/commit/da543ce14349e5f8f3eb3e8b6ee0df1ce284dac8))
* **mcp:** serialize message context result ([ebb1e07](https://github.com/verygoodplugins/whatsapp-mcp/commit/ebb1e07d8c22ccd62ac1b4dad7b67da7d63c50da))


### Documentation

* **readme:** document app state recovery ([#117](https://github.com/verygoodplugins/whatsapp-mcp/issues/117)) ([040f9e1](https://github.com/verygoodplugins/whatsapp-mcp/commit/040f9e11054fd0d5c1fc9e0ad690111bacca1447))

## [0.2.1](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.2.0...v0.2.1) (2026-05-14)


### Bug Fixes

* **bridge:** log send caller identity ([#96](https://github.com/verygoodplugins/whatsapp-mcp/issues/96)) ([c81947c](https://github.com/verygoodplugins/whatsapp-mcp/commit/c81947cbd4235e827637961b7ec53b47725d5a65))
* **bridge:** pass message store when sending messages ([#91](https://github.com/verygoodplugins/whatsapp-mcp/issues/91)) ([423ada9](https://github.com/verygoodplugins/whatsapp-mcp/commit/423ada9cd3432fcde79bb388d527a62cd5e8f696))
* **bridge:** set FileName and detect MIME for document sends ([#95](https://github.com/verygoodplugins/whatsapp-mcp/issues/95)) ([af18908](https://github.com/verygoodplugins/whatsapp-mcp/commit/af18908225767213bfedf970feb68f9a9bd79182))
* **ci:** make dependabot auto-approve non-fatal ([#65](https://github.com/verygoodplugins/whatsapp-mcp/issues/65)) ([ea0dfb1](https://github.com/verygoodplugins/whatsapp-mcp/commit/ea0dfb179ac0f68815e5a4b6b96058342590c8bd))
* send and track disappearing-message settings correctly ([#82](https://github.com/verygoodplugins/whatsapp-mcp/issues/82)) ([b250409](https://github.com/verygoodplugins/whatsapp-mcp/commit/b25040985b0e72d87d091d83d0bb64ff7d78eba8))
* **server:** list_chats/get_chat error when include_last_message=False ([#79](https://github.com/verygoodplugins/whatsapp-mcp/issues/79)) ([31e03e6](https://github.com/verygoodplugins/whatsapp-mcp/commit/31e03e63ddaa5ed2b0037faa7e70cb736addf617))


### Documentation

* add contributors section and updating instructions to README ([#67](https://github.com/verygoodplugins/whatsapp-mcp/issues/67)) ([71b0c03](https://github.com/verygoodplugins/whatsapp-mcp/commit/71b0c03070bb2054f45f202590c84d943cb2e4e6))
* Update SECURITY.md with enhanced security policy ([#78](https://github.com/verygoodplugins/whatsapp-mcp/issues/78)) ([214780f](https://github.com/verygoodplugins/whatsapp-mcp/commit/214780f3d1cd4c45f04676b37913c1d60b46f445))

## [0.2.0](https://github.com/verygoodplugins/whatsapp-mcp/compare/v0.1.0...v0.2.0) (2026-04-22)


### Features

* add image media support in webhook forwarding ([#45](https://github.com/verygoodplugins/whatsapp-mcp/issues/45)) ([c5f409c](https://github.com/verygoodplugins/whatsapp-mcp/commit/c5f409c78da8b7029b418e2fa428316dad7f970b))
* **bridge:** --full-history-pair flag to request full history at pair time ([#37](https://github.com/verygoodplugins/whatsapp-mcp/issues/37)) ([2358dd5](https://github.com/verygoodplugins/whatsapp-mcp/commit/2358dd5cf8cda17110076446c73637eab23bc7c3))
* **bridge:** capture incoming WhatsApp call events ([#39](https://github.com/verygoodplugins/whatsapp-mcp/issues/39)) ([197a0c9](https://github.com/verygoodplugins/whatsapp-mcp/commit/197a0c9de304b762db3e3a7fb98d9dd4ce510f74))


### Bug Fixes

* **bridge:** auto-download runs after StoreMessage to avoid lookup race ([#41](https://github.com/verygoodplugins/whatsapp-mcp/issues/41)) ([0c267f5](https://github.com/verygoodplugins/whatsapp-mcp/commit/0c267f5f0834138f54d7e2d731ffbd8f8deb5929))
* **bridge:** handle StreamReplaced event to recover from session conflicts ([#27](https://github.com/verygoodplugins/whatsapp-mcp/issues/27)) ([0cd6475](https://github.com/verygoodplugins/whatsapp-mcp/commit/0cd647560705acbbb3d46b236e1e09df7b20ceee))
* **bridge:** include message ID in media filenames to prevent same-second collisions ([#40](https://github.com/verygoodplugins/whatsapp-mcp/issues/40)) ([1e819c1](https://github.com/verygoodplugins/whatsapp-mcp/commit/1e819c1d67d7ed087ab3298253f4f987cd3e48e9))
* **bridge:** resolve [@lid](https://github.com/lid) sender to phone JID in webhook payload ([#56](https://github.com/verygoodplugins/whatsapp-mcp/issues/56)) ([746fca0](https://github.com/verygoodplugins/whatsapp-mcp/commit/746fca03297438f12b53ad3df4d33a6200773a05))
* **bridge:** surface image/video/document captions in extractTextContent ([#42](https://github.com/verygoodplugins/whatsapp-mcp/issues/42)) ([fbb3f28](https://github.com/verygoodplugins/whatsapp-mcp/commit/fbb3f283f0296a5f7e4aaf72eec2b99853de41b8))
* **mcp:** match messages by both phone number and LID via whatsmeow_lid_map ([#43](https://github.com/verygoodplugins/whatsapp-mcp/issues/43)) ([04c8755](https://github.com/verygoodplugins/whatsapp-mcp/commit/04c875568424a125b3e824ba094427f8d899d7c3))
* **mcp:** resolve contacts via whatsmeow store with LID → phone fallback ([#30](https://github.com/verygoodplugins/whatsapp-mcp/issues/30)) ([b9b0175](https://github.com/verygoodplugins/whatsapp-mcp/commit/b9b0175e6475d5402feedc190f44758045985992))
* pin `anyio<4.9` to avoid cancel scope regression ([#44](https://github.com/verygoodplugins/whatsapp-mcp/issues/44)) ([627db67](https://github.com/verygoodplugins/whatsapp-mcp/commit/627db6746076108ba4ca370ff3f95420ccbb30ef))
* security hardening for LAN exposure and Unicode search ([#55](https://github.com/verygoodplugins/whatsapp-mcp/issues/55)) ([8097f39](https://github.com/verygoodplugins/whatsapp-mcp/commit/8097f39dd4a19edd2ebab52d086704b528019005))


### Documentation

* add ROADMAP, AGENTS, CONTRIBUTING, CODEOWNERS, issue/PR templates ([#47](https://github.com/verygoodplugins/whatsapp-mcp/issues/47)) ([c2e4f5b](https://github.com/verygoodplugins/whatsapp-mcp/commit/c2e4f5b22cc3de42b4850222d3cc3ebfa4efbd14))
* document WHATSMEOW_DB_PATH env var ([#51](https://github.com/verygoodplugins/whatsapp-mcp/issues/51)) ([c24d151](https://github.com/verygoodplugins/whatsapp-mcp/commit/c24d151c9ebe07c052fa1e3a3662c10795747280))
* update remaining "go run main.go" references to "go run ." ([#50](https://github.com/verygoodplugins/whatsapp-mcp/issues/50)) ([53e1ca7](https://github.com/verygoodplugins/whatsapp-mcp/commit/53e1ca7dd2a470fd5f947f22ee667753cf7e6be6))

## 0.1.0 (2026-03-02)

### Added

**Go Bridge:**

- `/api/typing` endpoint - Send typing indicators to chats
- `/api/health` endpoint - Check WhatsApp connection status
- Webhook system for incoming messages (`webhook.go`)
  - Configurable via `WEBHOOK_URL` environment variable
  - Includes quoted message info (reply context)
  - `FORWARD_SELF` option to include self-sent messages
- Auto-download of media files when messages arrive
- Quoted message extraction (reply-to context: ID, sender, content)
- Connection status checks before API operations
- HTTP server timeouts for stability (read: 30s, write: 60s, idle: 120s)

**Python MCP Server:**

- `get_contact` tool - Resolve phone number to contact name
- `sender_display` field in messages - Shows "Name (phone)" format
- `sender_phone` field - Extracted phone number from JID
- Environment variable configuration (`WHATSAPP_DB_PATH`, `WHATSAPP_API_URL`)
- Test suite with pytest (8 tests)
- `sort_by` parameter for `list_messages` ("newest" or "oldest")
- Increased default limits (20 → 50 messages/chats)
- Maximum limits (500 messages, 200 chats)

**Infrastructure:**

- CI/CD pipeline with GitHub Actions (lint, build, test)
- Release Please workflow for automated release PRs/tags/changelog (`.github/workflows/release-please.yml`)
- Manual fallback release workflow for artifact re-publish (`.github/workflows/release.yml`)
- Version consistency check across `pyproject.toml`, `server.json`, and release tags
- Comprehensive documentation (README, CLAUDE.md)
- Maintainer release playbook (`docs/RELEASING.md`)
- Environment variable examples (`.env.example`)

### Fixed

- Go compilation errors from whatsmeow API changes (added `context.Background()`)
- Media filename consistency - now uses message timestamp instead of download time
- Startup migration now consolidates legacy `@lid` chat/message rows into mapped phone JIDs (`@s.whatsapp.net`) to prevent split chat history
- golangci.yml configuration (removed deprecated linters)
- Python linting issues (185 errors fixed)
- Trailing whitespace in SQL queries

### Changed

- Upgraded whatsmeow to latest version (v0.0.0-20260107124630)
- Upgraded Go linting pipeline to golangci-lint v2 (CI action + v2 config schema)
- Modernized Python type hints (`Optional[str]` → `str | None`)
- Improved docstrings with date format examples
- Media download uses message timestamp for consistent filenames
- Refactored `send_message` to use unified `recipient` parameter
- Adopted Release Please for automated versioning/changelog ([`#15`](https://github.com/verygoodplugins/whatsapp-mcp/issues/15)) ([6d45958](https://github.com/verygoodplugins/whatsapp-mcp/commit/6d45958139effa3079ff27a9708d400f89ba9ddf))

### Removed

- Unused `package.json` from Go bridge (was for wrangler, not needed)
- Compiled binary from git tracking

---

## Fork History

This project is a fork of [lharries/whatsapp-mcp](https://github.com/lharries/whatsapp-mcp), originally created by [Luke Harries](https://github.com/lharries) in 2024.

The original repository was last updated in April 2025. This fork by [Very Good Plugins](https://verygoodplugins.com) continues active development with bug fixes, new features, and MCP registry publication.
