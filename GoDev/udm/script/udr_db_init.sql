/*
 Navicat Premium Data Transfer

 Source Server         : 10.18.1.65mysql
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 10.18.1.65:3306
 Source Schema         : udr_db

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 21/12/2020 11:03:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access_mobility_subscription
-- ----------------------------
DROP TABLE IF EXISTS `access_mobility_subscription`;
CREATE TABLE `access_mobility_subscription`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `supi_id` int(10) UNSIGNED NOT NULL,
  `ue_ambr_uplink` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Pattern: \'^\\d+(\\.\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$\'\r\nExamples: \r\n\"125 Mbps\", \"0.125 Gbps\", \"125000 Kbps\"',
  `ue_ambr_downlink` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Pattern: \'^\\d+(\\.\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$\'',
  `nssai_id` int(10) UNSIGNED NOT NULL,
  `rfsp_index` smallint(5) UNSIGNED NOT NULL COMMENT '1-256\r\nIndex to RAT/Frequency Selection Priority',
  `subs_periodic_reg_timer` int(10) UNSIGNED NOT NULL COMMENT 'T3512, default 54 mins, second',
  `mps_priority_ind` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'True' COMMENT 'Indicates whether UE is subscribed to multimedia priority service',
  `mcs_priority_ind` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'False' COMMENT 'Indicates whether UE is subscribed to mission critical service',
  `active_time` int(10) UNSIGNED NOT NULL COMMENT 'subscribed active time for PSM UEs',
  `dl_packet_count` int(11) NOT NULL DEFAULT 0 COMMENT 'The following values are defined:\r\n0: \"Extended DL Data Buffering NOT REQUESTED\"\r\n-1: \"Extended DL Data Buffering REQUESTED, without a suggested number of packets\" \r\nn>0: \"Extended DL Data Buffering REQUESTED, with a suggested number of n packets\"',
  `mico_allowed` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'True' COMMENT 'Indicates whether the UE subscription allows MICO mode.',
  `supported_feature` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'pattern: [0-9][A-F]{0-16}',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `FK_ams_supi_id`(`supi_id`) USING BTREE,
  INDEX `fk_ams_nssai_id`(`nssai_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of access_mobility_subscription
-- ----------------------------
INSERT INTO `access_mobility_subscription` VALUES (1, 1, '200 Mbps', '200 Mbps', 1, 0, 3240, 'False', 'False', 0, 255, 'False', '8000');
INSERT INTO `access_mobility_subscription` VALUES (2, 2, '200 Mbps', '200 Mbps', 1, 0, 3240, 'False', 'False', 0, 255, 'False', '8000');

-- ----------------------------
-- Table structure for amf
-- ----------------------------
DROP TABLE IF EXISTS `amf`;
CREATE TABLE `amf`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `amf` varchar(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `amf`(`amf`) USING BTREE,
  INDEX `idx`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of amf
-- ----------------------------
INSERT INTO `amf` VALUES (1, '8000');

-- ----------------------------
-- Table structure for amf_3gpp_access_registration
-- ----------------------------
DROP TABLE IF EXISTS `amf_3gpp_access_registration`;
CREATE TABLE `amf_3gpp_access_registration`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `supi_id` int(11) NOT NULL,
  `amf_instance_id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `dereg_callback_uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `gam_amf_id` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `gam_plmn_id` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `rat_type` enum('NR','EUTAR','WLAN','VIRTUAL') CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `supported_features` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `purge_flag` tinyint(4) NULL DEFAULT NULL,
  `pei` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ims_vo_ps` enum('HOMOGENEOUS_SUPPORT','HOMOGENEOUS_NON_SUPPORT','NON_HOMOGENEOUS_OR_UNKNOWN') CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `amf_service_name_dereg` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pcscf_restoration_callback_uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `amf_service_name_pcscf_rest` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `initial_registration_ind` tinyint(4) NULL DEFAULT NULL,
  `dr_flag` tinyint(4) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of amf_3gpp_access_registration
-- ----------------------------
INSERT INTO `amf_3gpp_access_registration` VALUES (1, 1, '5105997e-b6c0-4d71-bd0f-f9e77013542b', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (2, 1, 'a23f4874-3b69-433c-bc94-57dfd9e4dbe2', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (3, 1, 'a23f4874-3b69-433c-bc94-57dfd9e4dbe2', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (4, 1, '1da4d743-f685-44d5-8213-2bb04d295934', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (5, 1, '1da4d743-f685-44d5-8213-2bb04d295934', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (6, 1, 'af137313-050b-4b87-917a-e58697ab5df7', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (7, 1, 'af137313-050b-4b87-917a-e58697ab5df7', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (8, 1, 'af137313-050b-4b87-917a-e58697ab5df7', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (9, 1, 'af137313-050b-4b87-917a-e58697ab5df7', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (10, 1, '270504af-48e8-4bbf-8342-ee5d9a60dd70', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (11, 1, '270504af-48e8-4bbf-8342-ee5d9a60dd70', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (12, 1, '270504af-48e8-4bbf-8342-ee5d9a60dd70', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (13, 1, '270504af-48e8-4bbf-8342-ee5d9a60dd70', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (14, 1, '270504af-48e8-4bbf-8342-ee5d9a60dd70', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (15, 1, 'dcbc732b-8556-4e7c-9826-04fbbfbd2c04', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (16, 1, 'fcedcc2a-f5d4-4227-a377-ae32541b8db2', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (17, 1, '506b2017-c81e-46ac-97cc-1b0a41aaa4f2', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (18, 1, '969ff2be-1f6c-41fa-aedc-a36b26f1783f', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (19, 1, '5cf36672-b1a9-475e-b260-03f49edb2b23', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (20, 1, 'd3a05ed5-4a67-47c4-bf4c-4aebf45bcd0a', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (21, 1, '9525ec98-efa0-40d2-9f7c-f6ef230cf08d', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (22, 1, 'e30d87ec-6887-4941-8f49-4f2046ab5593', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (23, 1, '8af716f9-21fe-4720-9a95-fa76221a2bc8', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (24, 1, '8af716f9-21fe-4720-9a95-fa76221a2bc8', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (25, 1, '8af716f9-21fe-4720-9a95-fa76221a2bc8', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (26, 1, '8af716f9-21fe-4720-9a95-fa76221a2bc8', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (27, 1, 'fc511b41-0494-470e-8096-363e6d9012c1', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (28, 1, 'cf55237b-9fa8-42cd-9e44-8c94fc01f6dd', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (29, 1, '163e2499-5154-4f1c-adc3-8f0bca58be31', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (30, 1, '587af459-7451-4163-a025-467e803ee445', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (31, 1, '587af459-7451-4163-a025-467e803ee445', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (32, 1, 'd8ce95d1-2d45-4f62-b549-3b6df54aa0cf', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (33, 1, 'a1bdb60a-82cc-41d1-9339-e8e6a544fccd', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (34, 1, 'e3298e19-3c2e-4123-9ab3-9664d92fc564', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (35, 1, 'e3298e19-3c2e-4123-9ab3-9664d92fc564', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (36, 1, 'e3298e19-3c2e-4123-9ab3-9664d92fc564', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);
INSERT INTO `amf_3gpp_access_registration` VALUES (37, 1, 'e3298e19-3c2e-4123-9ab3-9664d92fc564', 'http://10.18.1.52:29518/namf-comm/v1/nudm-uecm/imsi-460000234560001;guami=46000050081', '050081', '46000', 'NR', '', 0, '', 'HOMOGENEOUS_SUPPORT', '', '', '', 1, 0);

-- ----------------------------
-- Table structure for amf_backup_info
-- ----------------------------
DROP TABLE IF EXISTS `amf_backup_info`;
CREATE TABLE `amf_backup_info`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `amf_gpp_reg_id` bigint(20) NOT NULL,
  `back_up_amf` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `regid`(`amf_gpp_reg_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area`  (
  `id` int(10) UNSIGNED NOT NULL,
  `area_code` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'values are operator specific',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `area_code`(`area_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of area
-- ----------------------------
INSERT INTO `area` VALUES (1, '1');

-- ----------------------------
-- Table structure for arp
-- ----------------------------
DROP TABLE IF EXISTS `arp`;
CREATE TABLE `arp`  (
  `id` int(10) UNSIGNED NOT NULL,
  `priority_level` tinyint(3) UNSIGNED NOT NULL COMMENT 'ARP Priority Level, value [1-15] , 1 as the highest priority and 15 as the lowest priority',
  `preempt_cap` enum('MAY_PREEMPT','NOT_PREEMPT') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'NOT_PREEMPT' COMMENT 'Preemption Capability',
  `preempt_vuln` enum('PREEMPTABLE','NOT_PREEMPTABLE') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'PREEMPTABLE' COMMENT 'Preemption Vulnerability',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of arp
-- ----------------------------
INSERT INTO `arp` VALUES (1, 1, 'MAY_PREEMPT', 'NOT_PREEMPTABLE');

-- ----------------------------
-- Table structure for auth_data
-- ----------------------------
DROP TABLE IF EXISTS `auth_data`;
CREATE TABLE `auth_data`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `supi_id` int(10) UNSIGNED NOT NULL,
  `ki` varchar(33) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ki_k4_id` int(10) UNSIGNED NOT NULL,
  `op_id` int(10) UNSIGNED NOT NULL,
  `opc` varchar(33) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `opc_k4_id` int(10) UNSIGNED NOT NULL,
  `amf_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `fk_auth_data_1`(`supi_id`) USING BTREE,
  INDEX `fk_auth_data_2`(`op_id`) USING BTREE,
  INDEX `fk_auth_data_3`(`amf_id`) USING BTREE,
  INDEX `fk_auth_data_4`(`ki_k4_id`) USING BTREE,
  INDEX `fk_auth_data_5`(`opc_k4_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_data
-- ----------------------------
INSERT INTO `auth_data` VALUES (1, 1, '12345678901234567890123456789035', 1, 1, '', 1, 1);
INSERT INTO `auth_data` VALUES (2, 2, '12345678901234567890123456789035', 1, 1, '', 1, 1);

-- ----------------------------
-- Table structure for deregistration_data
-- ----------------------------
DROP TABLE IF EXISTS `deregistration_data`;
CREATE TABLE `deregistration_data`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `dereg_reason` enum('UE_INITIAL_REGISTRATION','UE_REGISTRATION_AREA_CHANGE','SUBSCRIPTION_WITHDRAWN','5GS_TO_EPS_MOBILITY','5GS_TO_EPS_MOBILITY_UE_INITIAL_REGISTRATION','REREGISTRATION_REQUIRED') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `access_type` enum('3GPP_ACCESS','NON_3GPP_ACCESS') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for dnn_info
-- ----------------------------
DROP TABLE IF EXISTS `dnn_info`;
CREATE TABLE `dnn_info`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `dnn` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `dnn_config_id` int(10) UNSIGNED NOT NULL,
  `default_dnn_ind` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'False' COMMENT 'Indicates whether this DNN is the default DNN:',
  `lbo_roaming_allowed` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'False' COMMENT 'indicates whether local breakout for the DNN is allowed when roaming',
  `iwk_eps_ind` enum('True','False') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'False' COMMENT 'Indicates whether interworking with EPS is subscribed:',
  `ladn_ind` enum('False','True') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'False' COMMENT 'Indicates whether the DNN is a local area data network.',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dnn`(`dnn`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `fk_dnninfo_configid`(`dnn_config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dnn_info
-- ----------------------------
INSERT INTO `dnn_info` VALUES (1, 'cmnet.com', 1, 'False', 'False', 'False', 'False');

-- ----------------------------
-- Table structure for dnn_configuration
-- ----------------------------
DROP TABLE IF EXISTS `dnn_configuration`;
CREATE TABLE `dnn_configuration`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `def_pdu_sess_type` enum('IPV4','IPV6','IPV4V6','ETHER','UNSTR') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'IPV4' COMMENT '\'ETHER\',\'UNSTR\',\'IPV6\',\'IPV4\',\'IPV4V6\'',
  `allowed_pdu_sess_type` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'IPV4' COMMENT 'Additional session types allowed for the data network\r\n\'ETHER\',\'UNSTR\',\'IPV6\',\'IPV4\',\'IPV4V6\'',
  `def_ssc_mode` enum('SSC_MODE_1','SSC_MODE_2','SSC_MODE_3') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'SSC_MODE_1' COMMENT '\'SSC_MODE_3\',\'SSC_MODE_2\',\'SSC_MODE_1\'',
  `allowed_ssc_mode` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'SSC_MODE_1' COMMENT '\'SSC_MODE_3\',\'SSC_MODE_2\',\'SSC_MODE_1\'',
  `iwk_eps_ind` enum('FALSE','TRUE') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'FALSE' COMMENT 'Indicates whether interworking with EPS is subscribed:',
  `ladn_ind` enum('FALSE','TRUE') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'FALSE' COMMENT 'Indicates whether the DNN is a local area data network',
  `5g_qos_profile_id` int(10) UNSIGNED NOT NULL,
  `sess_ambr_uplink` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '100 Mbps' COMMENT 'Pattern: \'^\\d+(\\.\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$\'\r\nExamples: \r\n\"125 Mbps\", \"0.125 Gbps\", \"125000 Kbps\"\r\n',
  `sess_ambr_downlink` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '100 Mbps' COMMENT 'Pattern: \'^\\d+(\\.\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$\'',
  `3gpp_charging_characteristic` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0000' COMMENT '[0-9][A-F]{4}',
  `up_security_integrity` enum('REQUIRED','PREFERRED','NOT_NEEDED') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'NOT_NEEDED',
  `up_security_confidentiality` enum('REQUIRED','PREFERRED','NOT_NEEDED') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'NOT_NEEDED',
  PRIMARY KEY (`id`, `3gpp_charging_characteristic`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `fk_dnnconfig_5g_qosprofile_id`(`5g_qos_profile_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dnn_configuration
-- ----------------------------
INSERT INTO `dnn_configuration` VALUES (1, 'cmnet.com', 'IPV4', 'IPV4', 'SSC_MODE_1', 'SSC_MODE_1', 'FALSE', 'FALSE', 1, '100 Mbps', '100 Mbps', '0000', 'REQUIRED', 'REQUIRED');

-- ----------------------------
-- Table structure for dynamic_uectxt_in_smf
-- ----------------------------
DROP TABLE IF EXISTS `dynamic_uectxt_in_smf`;
CREATE TABLE `dynamic_uectxt_in_smf`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `supi_id` int(10) UNSIGNED NOT NULL,
  `psi` tinyint(3) UNSIGNED NOT NULL,
  `dnn` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `smf_inst_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `plmnid` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'PlmnId: MNC + MCC\r\nPattern: \r\nMNC - \'^[0-9]{3}$\'\r\nMCC -  \'^[0-9]{2,3}$\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_pduSession_supi`(`supi_id`) USING BTREE,
  INDEX `FK_pduSession_plmnid`(`plmnid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dynamic_uectxt_in_smf
-- ----------------------------
INSERT INTO `dynamic_uectxt_in_smf` VALUES (1, 1, 5, 'cmnet.com', '', '46000');
INSERT INTO `dynamic_uectxt_in_smf` VALUES (2, 2, 5, 'cmnet.com', '', '46000');

-- ----------------------------
-- Table structure for gpsi
-- ----------------------------
DROP TABLE IF EXISTS `gpsi`;
CREATE TABLE `gpsi`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `gpsi` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'gpsi：either an External Id or an MSISDN. \r\nPattern：\'^(msisdn-[0-9]{5,15}|extid-.+@.+|.+)$\'',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `gpsi`(`gpsi`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gpsi
-- ----------------------------
INSERT INTO `gpsi` VALUES (1, '13712340001');
INSERT INTO `gpsi` VALUES (2, '13712340002');

-- ----------------------------
-- Table structure for guami
-- ----------------------------
DROP TABLE IF EXISTS `guami`;
CREATE TABLE `guami`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `amf_gpp_reg_id` bigint(20) NOT NULL,
  `amf_back_up_id` bigint(20) NOT NULL,
  `amf_id` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `plmn_id` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `bkid`(`amf_back_up_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for k4
-- ----------------------------
DROP TABLE IF EXISTS `k4`;
CREATE TABLE `k4`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(33) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of k4
-- ----------------------------
INSERT INTO `k4` VALUES (1, 'null', 'null');

-- ----------------------------
-- Table structure for map_area_tac
-- ----------------------------
DROP TABLE IF EXISTS `map_area_tac`;
CREATE TABLE `map_area_tac`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `area_id` int(10) UNSIGNED NOT NULL,
  `tac_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `area_id`(`area_id`) USING BTREE,
  INDEX `tac_id`(`tac_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for map_nssai_snssai
-- ----------------------------
DROP TABLE IF EXISTS `map_nssai_snssai`;
CREATE TABLE `map_nssai_snssai`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `nssai_id` int(10) UNSIGNED NOT NULL,
  `snssai_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_map_nssai_snssai_nid`(`nssai_id`) USING BTREE,
  INDEX `fk_map_nssai_snssai_sid`(`snssai_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of map_nssai_snssai
-- ----------------------------
INSERT INTO `map_nssai_snssai` VALUES (1, 1, 1);
INSERT INTO `map_nssai_snssai` VALUES (2, 1, 2);

-- ----------------------------
-- Table structure for map_supi_gpsi
-- ----------------------------
DROP TABLE IF EXISTS `map_supi_gpsi`;
CREATE TABLE `map_supi_gpsi`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'index pk',
  `supi_id` int(10) UNSIGNED NOT NULL,
  `gpsi_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `all`(`supi_id`, `gpsi_id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `gpsi_id`(`gpsi_id`) USING BTREE,
  INDEX `supi_id`(`supi_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of map_supi_gpsi
-- ----------------------------
INSERT INTO `map_supi_gpsi` VALUES (1, 1, 1);
INSERT INTO `map_supi_gpsi` VALUES (2, 2, 2);

-- ----------------------------
-- Table structure for nssai
-- ----------------------------
DROP TABLE IF EXISTS `nssai`;
CREATE TABLE `nssai`  (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `supported_feature` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'pattern: [0-9][A-F]{0-16}',
  `default_snssai_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  UNIQUE INDEX `default_snssai_id`(`default_snssai_id`, `name`, `supported_feature`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nssai
-- ----------------------------
INSERT INTO `nssai` VALUES (1, 'embb', '8000', 1);

-- ----------------------------
-- Table structure for op
-- ----------------------------
DROP TABLE IF EXISTS `op`;
CREATE TABLE `op`  (
  `id` int(10) UNSIGNED NOT NULL,
  `op` varchar(33) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `op`(`op`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of op
-- ----------------------------
INSERT INTO `op` VALUES (1, '12345678901234567890123456789012');

-- ----------------------------
-- Table structure for qos_profile_5g
-- ----------------------------
DROP TABLE IF EXISTS `qos_profile_5g`;
CREATE TABLE `qos_profile_5g`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `5qi` tinyint(3) UNSIGNED NOT NULL COMMENT ' 5G QoS Identifier， [0-255]',
  `priority_level` tinyint(3) UNSIGNED NOT NULL COMMENT '5QI Priority Level， [1-127] with 1 as the highest priority and 127 as the lowest priority',
  `arp_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  UNIQUE INDEX `all`(`5qi`, `priority_level`, `arp_id`) USING BTREE,
  INDEX `id`(`arp_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of qos_profile_5g
-- ----------------------------
INSERT INTO `qos_profile_5g` VALUES (1, 'qp_1', 9, 1, 1);

-- ----------------------------
-- Table structure for session_manage_info
-- ----------------------------
DROP TABLE IF EXISTS `session_manage_info`;
CREATE TABLE `session_manage_info`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `supi_id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `snssai_id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `dnn_id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `ue_dnn_config_id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `static_ipv4_address` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Pattern: \'^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$\'',
  `static_ipv6_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FK_smi_supi_id`(`supi_id`) USING BTREE,
  INDEX `FK_smi_dnn_id`(`dnn_id`) USING BTREE,
  INDEX `FK_smi_dnn_config_id`(`ue_dnn_config_id`) USING BTREE,
  INDEX `FK_smi_snssai_id`(`snssai_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of session_manage_info
-- ----------------------------
INSERT INTO `session_manage_info` VALUES (1, 0000000001, 0000000001, 0000000001, 0000000001, '10.55.1.1', '');

-- ----------------------------
-- Table structure for smf_registration
-- ----------------------------
DROP TABLE IF EXISTS `smf_registration`;
CREATE TABLE `smf_registration`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `supi` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `smf_instance_id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `supported_features` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pdu_session_id` int(10) UNSIGNED NOT NULL,
  `pcscf_restoration_callback_uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `plmn_id` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pgw_fqdn` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for smf_snssai
-- ----------------------------
DROP TABLE IF EXISTS `smf_snssai`;
CREATE TABLE `smf_snssai`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `smf_reg_id` bigint(20) UNSIGNED NOT NULL,
  `sst` int(10) UNSIGNED NULL DEFAULT NULL,
  `sd` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for snssai
-- ----------------------------
DROP TABLE IF EXISTS `snssai`;
CREATE TABLE `snssai`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `sst` tinyint(1) UNSIGNED ZEROFILL NOT NULL DEFAULT 0 COMMENT 'Slice/Service Type, 0-255',
  `sd` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Slice Differentiator, Pattern: \'^[A-Fa-f0-9]{6}$\'',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name-sst-sd`(`name`, `sst`, `sd`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of snssai
-- ----------------------------
INSERT INTO `snssai` VALUES (1, 'eMBB', 1, '');
INSERT INTO `snssai` VALUES (3, 'mMTC', 3, '');
INSERT INTO `snssai` VALUES (2, 'uRLLC', 2, '');
INSERT INTO `snssai` VALUES (4, 'V2X', 4, '');

-- ----------------------------
-- Table structure for snssai_info
-- ----------------------------
DROP TABLE IF EXISTS `snssai_info`;
CREATE TABLE `snssai_info`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `snssai_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `dnninfo_id` int(10) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `fk_sinfo_dnninfo_id`(`dnninfo_id`, `snssai_id`) USING BTREE,
  INDEX `fk_sinfo_snssai_id`(`snssai_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of snssai_info
-- ----------------------------
INSERT INTO `snssai_info` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for supi
-- ----------------------------
DROP TABLE IF EXISTS `supi`;
CREATE TABLE `supi`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'index, pk',
  `supi` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'supi： either an IMSI or an NAI\r\nPattern: \'^(imsi-[0-9]{5,15}|nai-.+|.+)$\'',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_supi`(`supi`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of supi
-- ----------------------------
INSERT INTO `supi` VALUES (1, '460000234560001');
INSERT INTO `supi` VALUES (2, '460000234560002');

-- ----------------------------
-- Table structure for tac
-- ----------------------------
DROP TABLE IF EXISTS `tac`;
CREATE TABLE `tac`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tac` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'TAC: [0-9][A-F]\r\nExamples:\r\nA legacy TAC 0x4305 shall be encoded as \"4305\".\r\nAn extended TAC 0x63F84B shall be encoded as \"63F84B\"\r\n',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `tac`(`tac`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tac
-- ----------------------------
INSERT INTO `tac` VALUES (1, '1');
INSERT INTO `tac` VALUES (2, '2');
INSERT INTO `tac` VALUES (3, '3');

-- ----------------------------
-- View structure for view_auth_data
-- ----------------------------
DROP VIEW IF EXISTS `view_auth_data`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_auth_data` AS select `supi`.`supi` AS `supi`,`auth_data`.`ki` AS `ki`,`auth_data`.`opc` AS `opc`,`op`.`op` AS `op`,`amf`.`amf` AS `amf` from ((((`auth_data` join `amf` on((`auth_data`.`amf_id` = `amf`.`id`))) join `op` on((`auth_data`.`op_id` = `op`.`id`))) join `supi` on((`auth_data`.`supi_id` = `supi`.`id`))) join `k4` on(((`auth_data`.`ki_k4_id` = `k4`.`id`) and (`auth_data`.`opc_k4_id` = `k4`.`id`))));

-- ----------------------------
-- View structure for view_dnn_configuration
-- ----------------------------
DROP VIEW IF EXISTS `view_dnn_configuration`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_dnn_configuration` AS select `dnn_configuration`.`id` AS `id`,`dnn_configuration`.`name` AS `name`,`dnn_configuration`.`def_pdu_sess_type` AS `def_pdu_sess_type`,`dnn_configuration`.`allowed_pdu_sess_type` AS `allowed_pdu_sess_type`,`dnn_configuration`.`def_ssc_mode` AS `def_ssc_mode`,`dnn_configuration`.`allowed_ssc_mode` AS `allowed_ssc_mode`,`dnn_configuration`.`iwk_eps_ind` AS `iwk_eps_ind`,`dnn_configuration`.`ladn_ind` AS `ladn_ind`,`dnn_configuration`.`sess_ambr_uplink` AS `sess_ambr_uplink`,`dnn_configuration`.`sess_ambr_downlink` AS `sess_ambr_downlink`,`dnn_configuration`.`3gpp_charging_characteristic` AS `3gpp_charging_characteristic`,`dnn_configuration`.`up_security_integrity` AS `up_security_integrity`,`dnn_configuration`.`up_security_confidentiality` AS `up_security_confidentiality`,`qos_profile_5g`.`5qi` AS `5qi`,`qos_profile_5g`.`priority_level` AS `priority_level`,`arp`.`priority_level` AS `arp_priority_level`,`arp`.`preempt_cap` AS `preempt_cap`,`arp`.`preempt_vuln` AS `preempt_vuln` from ((`dnn_configuration` join `qos_profile_5g` on((`dnn_configuration`.`5g_qos_profile_id` = `qos_profile_5g`.`id`))) join `arp` on((`qos_profile_5g`.`arp_id` = `arp`.`id`)));

-- ----------------------------
-- View structure for view_snssai_id_dnninfo
-- ----------------------------
DROP VIEW IF EXISTS `view_snssai_id_dnninfo`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_snssai_id_dnn_info` AS select `snssai_info`.`snssai_id` AS `snssai_id`,`dnn_info`.`dnn` AS `dnn`,`dnn_info`.`dnn_config_id` AS `dnn_config_id`,`dnn_info`.`default_dnn_ind` AS `default_dnn_ind`,`dnn_info`.`lbo_roaming_allowed` AS `lbo_roaming_allowed`,`dnn_info`.`iwk_eps_ind` AS `iwk_eps_ind`,`dnn_info`.`ladn_ind` AS `ladn_ind` from (`snssai_info` join `dnn_info` on((`snssai_info`.`dnninfo_id` = `dnn_info`.`id`)));

-- ----------------------------
-- View structure for view_supi_gpsi
-- ----------------------------
DROP VIEW IF EXISTS `view_supi_gpsi`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_supi_gpsi` AS select `supi`.`id` AS `id`,`supi`.`supi` AS `supi`,`gpsi`.`gpsi` AS `gpsi` from ((`supi` join `map_supi_gpsi` on((`map_supi_gpsi`.`supi_id` = `supi`.`id`))) join `gpsi` on((`map_supi_gpsi`.`gpsi_id` = `gpsi`.`id`)));

-- ----------------------------
-- View structure for view_supi_sm
-- ----------------------------
DROP VIEW IF EXISTS `view_supi_sm`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_supi_sm` AS select `supi`.`supi` AS `supi`,`session_manage_info`.`snssai_id` AS `snssai_id`,`dnn_info`.`dnn` AS `dnn`,`dnn_configuration`.`def_pdu_sess_type` AS `def_pdu_sess_type`,`dnn_configuration`.`allowed_pdu_sess_type` AS `allowed_pdu_sess_type`,`dnn_configuration`.`def_ssc_mode` AS `def_ssc_mode`,`dnn_configuration`.`allowed_ssc_mode` AS `allowed_ssc_mode`,`dnn_configuration`.`iwk_eps_ind` AS `iwk_eps_ind`,`dnn_configuration`.`ladn_ind` AS `ladn_ind`,`qos_profile_5g`.`5qi` AS `5qi`,`qos_profile_5g`.`priority_level` AS `priority_level`,`arp`.`priority_level` AS `arp_priority_level`,`arp`.`preempt_cap` AS `preempt_cap`,`arp`.`preempt_vuln` AS `preempt_vuln`,`dnn_configuration`.`sess_ambr_uplink` AS `sess_ambr_uplink`,`dnn_configuration`.`sess_ambr_downlink` AS `sess_ambr_downlink`,`dnn_configuration`.`3gpp_charging_characteristic` AS `3gpp_charging_characteristic`,`dnn_configuration`.`up_security_integrity` AS `up_security_integrity`,`dnn_configuration`.`up_security_confidentiality` AS `up_security_confidentiality`,`session_manage_info`.`static_ipv4_address` AS `static_ipv4_address`,`session_manage_info`.`static_ipv6_address` AS `static_ipv6_address` from (((((`session_manage_info` join `supi` on((`session_manage_info`.`supi_id` = `supi`.`id`))) join `dnn_configuration` on((`session_manage_info`.`ue_dnn_config_id` = `dnn_configuration`.`id`))) join `dnn_info` on((`session_manage_info`.`dnn_id` = `dnn_info`.`id`))) join `qos_profile_5g` on((`dnn_configuration`.`5g_qos_profile_id` = `qos_profile_5g`.`id`))) join `arp` on((`qos_profile_5g`.`arp_id` = `arp`.`id`)));

-- ----------------------------
-- View structure for view_supi_snssai_id
-- ----------------------------
DROP VIEW IF EXISTS `view_supi_snssai_id`;
CREATE ALGORITHM = UNDEFINED DEFINER = `5gc`@`%` SQL SECURITY DEFINER VIEW `view_supi_snssai_id` AS select `supi`.`id` AS `id`,`supi`.`supi` AS `supi`,`nssai`.`default_snssai_id` AS `default_snssai_id`,`nssai`.`supported_feature` AS `supported_feature`,`map_nssai_snssai`.`snssai_id` AS `snssai_id` from (((`supi` join `access_mobility_subscription` on((`access_mobility_subscription`.`supi_id` = `supi`.`id`))) join `nssai` on((`access_mobility_subscription`.`nssai_id` = `nssai`.`id`))) join `map_nssai_snssai` on((`map_nssai_snssai`.`nssai_id` = `nssai`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
