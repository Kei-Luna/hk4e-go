package cmd

import (
	"hk4e/protocol/proto"
)

const (
	AbilityChangeNotify                            = 1131
	AbilityInvocationFailNotify                    = 1107
	AbilityInvocationFixedNotify                   = 1172
	AbilityInvocationsNotify                       = 1198
	AcceptCityReputationRequestReq                 = 2890
	AcceptCityReputationRequestRsp                 = 2873
	AchievementAllDataNotify                       = 2676
	AchievementUpdateNotify                        = 2668
	ActivityAcceptAllGiveGiftReq                   = 8113
	ActivityAcceptAllGiveGiftRsp                   = 8132
	ActivityAcceptGiveGiftReq                      = 8095
	ActivityAcceptGiveGiftRsp                      = 8502
	ActivityBannerClearReq                         = 2009
	ActivityBannerClearRsp                         = 2163
	ActivityBannerNotify                           = 2155
	ActivityCoinInfoNotify                         = 2008
	ActivityCondStateChangeNotify                  = 2140
	ActivityDisableTransferPointInteractionNotify  = 8982
	ActivityGetCanGiveFriendGiftReq                = 8559
	ActivityGetCanGiveFriendGiftRsp                = 8848
	ActivityGetFriendGiftWishListReq               = 8806
	ActivityGetFriendGiftWishListRsp               = 8253
	ActivityGetRecvGiftListReq                     = 8725
	ActivityGetRecvGiftListRsp                     = 8120
	ActivityGiveFriendGiftReq                      = 8233
	ActivityGiveFriendGiftRsp                      = 8696
	ActivityHaveRecvGiftNotify                     = 8733
	ActivityInfoNotify                             = 2060
	ActivityPlayOpenAnimNotify                     = 2157
	ActivityPushTipsInfoNotify                     = 8513
	ActivityReadPushTipsReq                        = 8145
	ActivityReadPushTipsRsp                        = 8574
	ActivitySaleChangeNotify                       = 2071
	ActivityScheduleInfoNotify                     = 2073
	ActivitySelectAvatarCardReq                    = 2028
	ActivitySelectAvatarCardRsp                    = 2189
	ActivitySetGiftWishReq                         = 8017
	ActivitySetGiftWishRsp                         = 8554
	ActivityTakeAllScoreRewardReq                  = 8372
	ActivityTakeAllScoreRewardRsp                  = 8043
	ActivityTakeScoreRewardReq                     = 8971
	ActivityTakeScoreRewardRsp                     = 8583
	ActivityTakeWatcherRewardBatchReq              = 2159
	ActivityTakeWatcherRewardBatchRsp              = 2109
	ActivityTakeWatcherRewardReq                   = 2038
	ActivityTakeWatcherRewardRsp                   = 2034
	ActivityUpdateWatcherNotify                    = 2156
	AddAranaraCollectionNotify                     = 6368
	AddBackupAvatarTeamReq                         = 1687
	AddBackupAvatarTeamRsp                         = 1735
	AddBlacklistReq                                = 4088
	AddBlacklistRsp                                = 4026
	AddFriendNotify                                = 4022
	AddNoGachaAvatarCardNotify                     = 1655
	AddQuestContentProgressReq                     = 421
	AddQuestContentProgressRsp                     = 403
	AddRandTaskInfoNotify                          = 119
	AddSeenMonsterNotify                           = 223
	AdjustWorldLevelReq                            = 164
	AdjustWorldLevelRsp                            = 138
	AllCoopInfoNotify                              = 1976
	AllMarkPointNotify                             = 3283
	AllSeenMonsterNotify                           = 271
	AllShareCDDataNotify                           = 9072
	AllWidgetBackgroundActiveStateNotify           = 6092
	AllWidgetDataNotify                            = 4271
	AnchorPointDataNotify                          = 4276
	AnchorPointOpReq                               = 4257
	AnchorPointOpRsp                               = 4252
	AnimatorForceSetAirMoveNotify                  = 374
	AntiAddictNotify                               = 180
	AranaraCollectionDataNotify                    = 6376
	AreaPlayInfoNotify                             = 3323
	ArenaChallengeFinishNotify                     = 2030
	AskAddFriendNotify                             = 4065
	AskAddFriendReq                                = 4007
	AskAddFriendRsp                                = 4021
	AssociateInferenceWordReq                      = 429
	AssociateInferenceWordRsp                      = 457
	AsterLargeInfoNotify                           = 2146
	AsterLittleInfoNotify                          = 2068
	AsterMidCampInfoNotify                         = 2133
	AsterMidInfoNotify                             = 2031
	AsterMiscInfoNotify                            = 2036
	AsterProgressInfoNotify                        = 2016
	AvatarAddNotify                                = 1769
	AvatarBuffAddNotify                            = 388
	AvatarBuffDelNotify                            = 326
	AvatarCardChangeReq                            = 688
	AvatarCardChangeRsp                            = 626
	AvatarChangeAnimHashReq                        = 1711
	AvatarChangeAnimHashRsp                        = 1647
	AvatarChangeCostumeNotify                      = 1644
	AvatarChangeCostumeReq                         = 1778
	AvatarChangeCostumeRsp                         = 1645
	AvatarChangeElementTypeReq                     = 1785
	AvatarChangeElementTypeRsp                     = 1651
	AvatarDataNotify                               = 1633
	AvatarDelNotify                                = 1773
	AvatarDieAnimationEndReq                       = 1610
	AvatarDieAnimationEndRsp                       = 1694
	AvatarEnterElementViewNotify                   = 334
	AvatarEquipAffixStartNotify                    = 1662
	AvatarEquipChangeNotify                        = 647
	AvatarExpeditionAllDataReq                     = 1722
	AvatarExpeditionAllDataRsp                     = 1648
	AvatarExpeditionCallBackReq                    = 1752
	AvatarExpeditionCallBackRsp                    = 1726
	AvatarExpeditionDataNotify                     = 1771
	AvatarExpeditionGetRewardReq                   = 1623
	AvatarExpeditionGetRewardRsp                   = 1784
	AvatarExpeditionStartReq                       = 1715
	AvatarExpeditionStartRsp                       = 1719
	AvatarFetterDataNotify                         = 1782
	AvatarFetterLevelRewardReq                     = 1653
	AvatarFetterLevelRewardRsp                     = 1606
	AvatarFightPropNotify                          = 1207
	AvatarFightPropUpdateNotify                    = 1221
	AvatarFlycloakChangeNotify                     = 1643
	AvatarFollowRouteNotify                        = 3458
	AvatarGainCostumeNotify                        = 1677
	AvatarGainFlycloakNotify                       = 1656
	AvatarLifeStateChangeNotify                    = 1290
	AvatarPromoteGetRewardReq                      = 1696
	AvatarPromoteGetRewardRsp                      = 1683
	AvatarPromoteReq                               = 1664
	AvatarPromoteRsp                               = 1639
	AvatarPropChangeReasonNotify                   = 1273
	AvatarPropNotify                               = 1231
	AvatarSatiationDataNotify                      = 1693
	AvatarSkillChangeNotify                        = 1097
	AvatarSkillDepotChangeNotify                   = 1035
	AvatarSkillInfoNotify                          = 1090
	AvatarSkillMaxChargeCountNotify                = 1003
	AvatarSkillUpgradeReq                          = 1075
	AvatarSkillUpgradeRsp                          = 1048
	AvatarTeamAllDataNotify                        = 1749
	AvatarTeamUpdateNotify                         = 1706
	AvatarUnlockTalentNotify                       = 1012
	AvatarUpgradeReq                               = 1770
	AvatarUpgradeRsp                               = 1701
	AvatarWearFlycloakReq                          = 1737
	AvatarWearFlycloakRsp                          = 1698
	BackMyWorldReq                                 = 286
	BackMyWorldRsp                                 = 201
	BackPlayCustomDungeonOfficialReq               = 6203
	BackPlayCustomDungeonOfficialRsp               = 6204
	BackRebornGalleryReq                           = 5593
	BackRebornGalleryRsp                           = 5527
	BargainOfferPriceReq                           = 493
	BargainOfferPriceRsp                           = 427
	BargainStartNotify                             = 404
	BargainTerminateNotify                         = 494
	BartenderCancelLevelReq                        = 8771
	BartenderCancelLevelRsp                        = 8686
	BartenderCancelOrderReq                        = 8442
	BartenderCancelOrderRsp                        = 8837
	BartenderCompleteOrderReq                      = 8414
	BartenderCompleteOrderRsp                      = 8125
	BartenderFinishLevelReq                        = 8227
	BartenderFinishLevelRsp                        = 8093
	BartenderGetFormulaReq                         = 8462
	BartenderGetFormulaRsp                         = 8842
	BartenderLevelProgressNotify                   = 8756
	BartenderStartLevelReq                         = 8507
	BartenderStartLevelRsp                         = 8402
	BattlePassAllDataNotify                        = 2626
	BattlePassBuySuccNotify                        = 2614
	BattlePassCurScheduleUpdateNotify              = 2607
	BattlePassMissionDelNotify                     = 2625
	BattlePassMissionUpdateNotify                  = 2618
	BeginCameraSceneLookNotify                     = 270
	BeginCameraSceneLookWithTemplateNotify         = 3160
	BigTalentPointConvertReq                       = 1007
	BigTalentPointConvertRsp                       = 1021
	BlessingAcceptAllGivePicReq                    = 2045
	BlessingAcceptAllGivePicRsp                    = 2044
	BlessingAcceptGivePicReq                       = 2006
	BlessingAcceptGivePicRsp                       = 2055
	BlessingGetAllRecvPicRecordListReq             = 2096
	BlessingGetAllRecvPicRecordListRsp             = 2083
	BlessingGetFriendPicListReq                    = 2043
	BlessingGetFriendPicListRsp                    = 2056
	BlessingGiveFriendPicReq                       = 2062
	BlessingGiveFriendPicRsp                       = 2053
	BlessingRecvFriendPicNotify                    = 2178
	BlessingRedeemRewardReq                        = 2137
	BlessingRedeemRewardRsp                        = 2098
	BlessingScanReq                                = 2081
	BlessingScanRsp                                = 2093
	BlitzRushParkourRestartReq                     = 8653
	BlitzRushParkourRestartRsp                     = 8944
	BlossomBriefInfoNotify                         = 2712
	BlossomChestCreateNotify                       = 2721
	BlossomChestInfoNotify                         = 890
	BonusActivityInfoReq                           = 2548
	BonusActivityInfoRsp                           = 2597
	BonusActivityUpdateNotify                      = 2575
	BossChestActivateNotify                        = 803
	BounceConjuringSettleNotify                    = 8084
	BuoyantCombatSettleNotify                      = 8305
	BuyBattlePassLevelReq                          = 2647
	BuyBattlePassLevelRsp                          = 2637
	BuyGoodsReq                                    = 712
	BuyGoodsRsp                                    = 735
	BuyResinReq                                    = 602
	BuyResinRsp                                    = 619
	CalcWeaponUpgradeReturnItemsReq                = 633
	CalcWeaponUpgradeReturnItemsRsp                = 684
	CanUseSkillNotify                              = 1005
	CancelCityReputationRequestReq                 = 2899
	CancelCityReputationRequestRsp                 = 2831
	CancelCoopTaskReq                              = 1997
	CancelCoopTaskRsp                              = 1987
	CancelFinishParentQuestNotify                  = 424
	CardProductRewardNotify                        = 4107
	CataLogFinishedGlobalWatcherAllDataNotify      = 6370
	CataLogNewFinishedGlobalWatcherNotify          = 6395
	ChallengeDataNotify                            = 953
	ChallengeRecordNotify                          = 993
	ChangeAvatarReq                                = 1640
	ChangeAvatarRsp                                = 1607
	ChangeCustomDungeonRoomReq                     = 6222
	ChangeCustomDungeonRoomRsp                     = 6244
	ChangeGameTimeReq                              = 173
	ChangeGameTimeRsp                              = 199
	ChangeMailStarNotify                           = 1448
	ChangeMpTeamAvatarReq                          = 1708
	ChangeMpTeamAvatarRsp                          = 1753
	ChangeServerGlobalValueNotify                  = 27
	ChangeTeamNameReq                              = 1603
	ChangeTeamNameRsp                              = 1666
	ChangeWidgetBackgroundActiveStateReq           = 5907
	ChangeWidgetBackgroundActiveStateRsp           = 6060
	ChangeWorldToSingleModeNotify                  = 3006
	ChangeWorldToSingleModeReq                     = 3066
	ChangeWorldToSingleModeRsp                     = 3282
	ChannelerSlabCheckEnterLoopDungeonReq          = 8745
	ChannelerSlabCheckEnterLoopDungeonRsp          = 8452
	ChannelerSlabEnterLoopDungeonReq               = 8869
	ChannelerSlabEnterLoopDungeonRsp               = 8081
	ChannelerSlabLoopDungeonChallengeInfoNotify    = 8224
	ChannelerSlabLoopDungeonSelectConditionReq     = 8503
	ChannelerSlabLoopDungeonSelectConditionRsp     = 8509
	ChannelerSlabLoopDungeonTakeFirstPassRewardReq = 8589
	ChannelerSlabLoopDungeonTakeFirstPassRewardRsp = 8539
	ChannelerSlabLoopDungeonTakeScoreRewardReq     = 8684
	ChannelerSlabLoopDungeonTakeScoreRewardRsp     = 8433
	ChannelerSlabOneOffDungeonInfoNotify           = 8729
	ChannelerSlabOneOffDungeonInfoReq              = 8409
	ChannelerSlabOneOffDungeonInfoRsp              = 8268
	ChannelerSlabSaveAssistInfoReq                 = 8416
	ChannelerSlabSaveAssistInfoRsp                 = 8932
	ChannelerSlabStageActiveChallengeIndexNotify   = 8734
	ChannelerSlabStageOneofDungeonNotify           = 8203
	ChannelerSlabTakeoffBuffReq                    = 8516
	ChannelerSlabTakeoffBuffRsp                    = 8237
	ChannelerSlabWearBuffReq                       = 8107
	ChannelerSlabWearBuffRsp                       = 8600
	ChapterStateNotify                             = 405
	CharAmusementSettleNotify                      = 23133
	ChatChannelDataNotify                          = 4998
	ChatChannelUpdateNotify                        = 5025
	ChatHistoryNotify                              = 3496
	CheckAddItemExceedLimitNotify                  = 692
	CheckGroupReplacedReq                          = 3113
	CheckGroupReplacedRsp                          = 3152
	CheckSegmentCRCNotify                          = 39
	CheckSegmentCRCReq                             = 53
	CheckUgcStateReq                               = 6342
	CheckUgcStateRsp                               = 6314
	CheckUgcUpdateReq                              = 6320
	CheckUgcUpdateRsp                              = 6345
	ChessEscapedMonstersNotify                     = 5314
	ChessLeftMonstersNotify                        = 5360
	ChessManualRefreshCardsReq                     = 5389
	ChessManualRefreshCardsRsp                     = 5359
	ChessPickCardNotify                            = 5380
	ChessPickCardReq                               = 5333
	ChessPickCardRsp                               = 5384
	ChessPlayerInfoNotify                          = 5332
	ChessSelectedCardsNotify                       = 5392
	ChooseCurAvatarTeamReq                         = 1796
	ChooseCurAvatarTeamRsp                         = 1661
	CityReputationDataNotify                       = 2805
	CityReputationLevelupNotify                    = 2807
	ClearRoguelikeCurseNotify                      = 8207
	ClientAIStateNotify                            = 1181
	ClientAbilitiesInitFinishCombineNotify         = 1103
	ClientAbilityChangeNotify                      = 1175
	ClientAbilityInitBeginNotify                   = 1112
	ClientAbilityInitFinishNotify                  = 1135
	ClientBulletCreateNotify                       = 4
	ClientCollectorDataNotify                      = 4264
	ClientHashDebugNotify                          = 3086
	ClientLoadingCostumeVerificationNotify         = 3487
	ClientLockGameTimeNotify                       = 114
	ClientNewMailNotify                            = 1499
	ClientPauseNotify                              = 260
	ClientReconnectNotify                          = 75
	ClientRemoveCombatEndModifierNotify            = 1182
	ClientReportNotify                             = 81
	ClientScriptEventNotify                        = 213
	ClientTransmitReq                              = 291
	ClientTransmitRsp                              = 224
	ClientTriggerEventNotify                       = 148
	CloseCommonTipsNotify                          = 3194
	ClosedItemNotify                               = 614
	CodexDataFullNotify                            = 4205
	CodexDataUpdateNotify                          = 4207
	CombatInvocationsNotify                        = 319
	CombineDataNotify                              = 659
	CombineFormulaDataNotify                       = 632
	CombineReq                                     = 643
	CombineRsp                                     = 674
	CommonPlayerTipsNotify                         = 8466
	CompoundDataNotify                             = 146
	CompoundUnlockNotify                           = 128
	CookDataNotify                                 = 195
	CookGradeDataNotify                            = 134
	CookRecipeDataNotify                           = 106
	CoopCgShowNotify                               = 1983
	CoopCgUpdateNotify                             = 1994
	CoopChapterUpdateNotify                        = 1972
	CoopDataNotify                                 = 1979
	CoopPointUpdateNotify                          = 1991
	CoopProgressUpdateNotify                       = 1998
	CoopRewardUpdateNotify                         = 1999
	CreateMassiveEntityNotify                      = 367
	CreateMassiveEntityReq                         = 342
	CreateMassiveEntityRsp                         = 330
	CreateVehicleReq                               = 893
	CreateVehicleRsp                               = 827
	CrystalLinkDungeonInfoNotify                   = 8858
	CrystalLinkEnterDungeonReq                     = 8325
	CrystalLinkEnterDungeonRsp                     = 8147
	CrystalLinkRestartDungeonReq                   = 8022
	CrystalLinkRestartDungeonRsp                   = 8119
	CustomDungeonBattleRecordNotify                = 6236
	CustomDungeonOfficialNotify                    = 6221
	CustomDungeonRecoverNotify                     = 6217
	CustomDungeonUpdateNotify                      = 6223
	CutSceneBeginNotify                            = 296
	CutSceneEndNotify                              = 215
	CutSceneFinishNotify                           = 262
	DailyTaskDataNotify                            = 158
	DailyTaskFilterCityReq                         = 111
	DailyTaskFilterCityRsp                         = 144
	DailyTaskProgressNotify                        = 170
	DailyTaskScoreRewardNotify                     = 117
	DailyTaskUnlockedCitiesNotify                  = 186
	DataResVersionNotify                           = 167
	DealAddFriendReq                               = 4003
	DealAddFriendRsp                               = 4090
	DeathZoneInfoNotify                            = 6268
	DeathZoneObserveNotify                         = 3475
	DebugNotify                                    = 101
	DelBackupAvatarTeamReq                         = 1731
	DelBackupAvatarTeamRsp                         = 1729
	DelMailReq                                     = 1421
	DelMailRsp                                     = 1403
	DelScenePlayTeamEntityNotify                   = 3318
	DelTeamEntityNotify                            = 302
	DeleteFriendNotify                             = 4053
	DeleteFriendReq                                = 4031
	DeleteFriendRsp                                = 4075
	DeshretObeliskChestInfoNotify                  = 841
	DestroyMassiveEntityNotify                     = 358
	DestroyMaterialReq                             = 640
	DestroyMaterialRsp                             = 618
	DigActivityChangeGadgetStateReq                = 8464
	DigActivityChangeGadgetStateRsp                = 8430
	DigActivityMarkPointChangeNotify               = 8109
	DisableRoguelikeTrapNotify                     = 8259
	DoGachaReq                                     = 1512
	DoGachaRsp                                     = 1535
	DoRoguelikeDungeonCardGachaReq                 = 8148
	DoRoguelikeDungeonCardGachaRsp                 = 8472
	DoSetPlayerBornDataNotify                      = 147
	DraftGuestReplyInviteNotify                    = 5490
	DraftGuestReplyInviteReq                       = 5421
	DraftGuestReplyInviteRsp                       = 5403
	DraftGuestReplyTwiceConfirmNotify              = 5497
	DraftGuestReplyTwiceConfirmReq                 = 5431
	DraftGuestReplyTwiceConfirmRsp                 = 5475
	DraftInviteResultNotify                        = 5473
	DraftOwnerInviteNotify                         = 5407
	DraftOwnerStartInviteReq                       = 5412
	DraftOwnerStartInviteRsp                       = 5435
	DraftOwnerTwiceConfirmNotify                   = 5499
	DraftTwiceConfirmResultNotify                  = 5448
	DragonSpineChapterFinishNotify                 = 2069
	DragonSpineChapterOpenNotify                   = 2022
	DragonSpineChapterProgressChangeNotify         = 2065
	DragonSpineCoinChangeNotify                    = 2088
	DropHintNotify                                 = 650
	DropItemReq                                    = 699
	DropItemRsp                                    = 631
	DungeonCandidateTeamChangeAvatarReq            = 956
	DungeonCandidateTeamChangeAvatarRsp            = 942
	DungeonCandidateTeamCreateReq                  = 995
	DungeonCandidateTeamCreateRsp                  = 906
	DungeonCandidateTeamDismissNotify              = 963
	DungeonCandidateTeamInfoNotify                 = 927
	DungeonCandidateTeamInviteNotify               = 994
	DungeonCandidateTeamInviteReq                  = 934
	DungeonCandidateTeamInviteRsp                  = 950
	DungeonCandidateTeamKickReq                    = 943
	DungeonCandidateTeamKickRsp                    = 974
	DungeonCandidateTeamLeaveReq                   = 976
	DungeonCandidateTeamLeaveRsp                   = 946
	DungeonCandidateTeamPlayerLeaveNotify          = 926
	DungeonCandidateTeamRefuseNotify               = 988
	DungeonCandidateTeamReplyInviteReq             = 941
	DungeonCandidateTeamReplyInviteRsp             = 949
	DungeonCandidateTeamSetChangingAvatarReq       = 918
	DungeonCandidateTeamSetChangingAvatarRsp       = 966
	DungeonCandidateTeamSetReadyReq                = 991
	DungeonCandidateTeamSetReadyRsp                = 924
	DungeonChallengeBeginNotify                    = 947
	DungeonChallengeFinishNotify                   = 939
	DungeonDataNotify                              = 982
	DungeonDieOptionReq                            = 975
	DungeonDieOptionRsp                            = 948
	DungeonEntryInfoReq                            = 972
	DungeonEntryInfoRsp                            = 998
	DungeonEntryToBeExploreNotify                  = 3147
	DungeonFollowNotify                            = 922
	DungeonGetStatueDropReq                        = 965
	DungeonGetStatueDropRsp                        = 904
	DungeonInterruptChallengeReq                   = 917
	DungeonInterruptChallengeRsp                   = 902
	DungeonPlayerDieNotify                         = 931
	DungeonPlayerDieReq                            = 981
	DungeonPlayerDieRsp                            = 905
	DungeonRestartInviteNotify                     = 957
	DungeonRestartInviteReplyNotify                = 987
	DungeonRestartInviteReplyReq                   = 1000
	DungeonRestartInviteReplyRsp                   = 916
	DungeonRestartReq                              = 961
	DungeonRestartResultNotify                     = 940
	DungeonRestartRsp                              = 929
	DungeonReviseLevelNotify                       = 933
	DungeonSettleNotify                            = 999
	DungeonShowReminderNotify                      = 997
	DungeonSlipRevivePointActivateReq              = 958
	DungeonSlipRevivePointActivateRsp              = 970
	DungeonWayPointActivateReq                     = 990
	DungeonWayPointActivateRsp                     = 973
	DungeonWayPointNotify                          = 903
	EchoNotify                                     = 65
	EchoShellTakeRewardReq                         = 8114
	EchoShellTakeRewardRsp                         = 8797
	EchoShellUpdateNotify                          = 8150
	EffigyChallengeInfoNotify                      = 2090
	EffigyChallengeResultNotify                    = 2046
	EffigyChallengeV2ChooseSkillReq                = 21269
	EffigyChallengeV2ChooseSkillRsp                = 22448
	EffigyChallengeV2DungeonInfoNotify             = 22835
	EffigyChallengeV2EnterDungeonReq               = 23489
	EffigyChallengeV2EnterDungeonRsp               = 24917
	EffigyChallengeV2RestartDungeonReq             = 24522
	EffigyChallengeV2RestartDungeonRsp             = 23167
	EndCameraSceneLookNotify                       = 217
	EnterChessDungeonReq                           = 8191
	EnterChessDungeonRsp                           = 8592
	EnterCustomDungeonReq                          = 6226
	EnterCustomDungeonRsp                          = 6218
	EnterFishingReq                                = 5826
	EnterFishingRsp                                = 5818
	EnterFungusFighterPlotDungeonReq               = 23053
	EnterFungusFighterPlotDungeonRsp               = 21008
	EnterFungusFighterTrainingDungeonReq           = 23860
	EnterFungusFighterTrainingDungeonRsp           = 21593
	EnterIrodoriChessDungeonReq                    = 8717
	EnterIrodoriChessDungeonRsp                    = 8546
	EnterMechanicusDungeonReq                      = 3931
	EnterMechanicusDungeonRsp                      = 3975
	EnterRogueDiaryDungeonReq                      = 8943
	EnterRogueDiaryDungeonRsp                      = 8352
	EnterRoguelikeDungeonNotify                    = 8652
	EnterSceneDoneReq                              = 277
	EnterSceneDoneRsp                              = 237
	EnterScenePeerNotify                           = 252
	EnterSceneReadyReq                             = 208
	EnterSceneReadyRsp                             = 209
	EnterSceneWeatherAreaNotify                    = 256
	EnterTransPointRegionNotify                    = 205
	EnterTrialAvatarActivityDungeonReq             = 2118
	EnterTrialAvatarActivityDungeonRsp             = 2183
	EnterWorldAreaReq                              = 250
	EnterWorldAreaRsp                              = 243
	EntityAiKillSelfNotify                         = 340
	EntityAiSyncNotify                             = 400
	EntityAuthorityChangeNotify                    = 394
	EntityConfigHashNotify                         = 3189
	EntityFightPropChangeReasonNotify              = 1203
	EntityFightPropNotify                          = 1212
	EntityFightPropUpdateNotify                    = 1235
	EntityForceSyncReq                             = 274
	EntityForceSyncRsp                             = 276
	EntityJumpNotify                               = 222
	EntityMoveRoomNotify                           = 3178
	EntityPropNotify                               = 1272
	EntityTagChangeNotify                          = 3316
	EquipRoguelikeRuneReq                          = 8306
	EquipRoguelikeRuneRsp                          = 8705
	EvtAiSyncCombatThreatInfoNotify                = 329
	EvtAiSyncSkillCdNotify                         = 376
	EvtAnimatorParameterNotify                     = 398
	EvtAnimatorStateChangedNotify                  = 331
	EvtAvatarEnterFocusNotify                      = 304
	EvtAvatarExitFocusNotify                       = 393
	EvtAvatarLockChairReq                          = 318
	EvtAvatarLockChairRsp                          = 366
	EvtAvatarSitDownNotify                         = 324
	EvtAvatarStandUpNotify                         = 356
	EvtAvatarUpdateFocusNotify                     = 327
	EvtBeingHealedNotify                           = 333
	EvtBeingHitNotify                              = 372
	EvtBeingHitsCombineNotify                      = 346
	EvtBulletDeactiveNotify                        = 397
	EvtBulletHitNotify                             = 348
	EvtBulletMoveNotify                            = 365
	EvtCostStaminaNotify                           = 373
	EvtCreateGadgetNotify                          = 307
	EvtDestroyGadgetNotify                         = 321
	EvtDestroyServerGadgetNotify                   = 387
	EvtDoSkillSuccNotify                           = 335
	EvtEntityRenderersChangedNotify                = 343
	EvtEntityStartDieEndNotify                     = 381
	EvtFaceToDirNotify                             = 390
	EvtFaceToEntityNotify                          = 303
	EvtLocalGadgetOwnerLeaveSceneNotify            = 384
	EvtRushMoveNotify                              = 375
	EvtSetAttackTargetNotify                       = 399
	ExclusiveRuleNotify                            = 101
	ExecuteGadgetLuaReq                            = 269
	ExecuteGadgetLuaRsp                            = 210
	ExecuteGroupTriggerReq                         = 257
	ExecuteGroupTriggerRsp                         = 300
	ExitCustomDungeonTryReq                        = 6247
	ExitCustomDungeonTryRsp                        = 6237
	ExitFishingReq                                 = 5814
	ExitFishingRsp                                 = 5847
	ExitSceneWeatherAreaNotify                     = 242
	ExitTransPointRegionNotify                     = 282
	ExpeditionChallengeEnterRegionNotify           = 2154
	ExpeditionChallengeFinishedNotify              = 2091
	ExpeditionRecallReq                            = 2131
	ExpeditionRecallRsp                            = 2129
	ExpeditionStartReq                             = 2087
	ExpeditionStartRsp                             = 2135
	ExpeditionTakeRewardReq                        = 2149
	ExpeditionTakeRewardRsp                        = 2080
	FindHilichurlAcceptQuestNotify                 = 8659
	FindHilichurlFinishSecondQuestNotify           = 8901
	FinishDeliveryNotify                           = 2089
	FinishLanternProjectionReq                     = 8704
	FinishLanternProjectionRsp                     = 8713
	FinishMainCoopReq                              = 1952
	FinishMainCoopRsp                              = 1981
	FinishedParentQuestNotify                      = 435
	FinishedParentQuestUpdateNotify                = 407
	FinishedTalkIdListNotify                       = 573
	FireworksLaunchDataNotify                      = 5928
	FireworksReformDataNotify                      = 6033
	FishAttractNotify                              = 5837
	FishBaitGoneNotify                             = 5823
	FishBattleBeginReq                             = 5820
	FishBattleBeginRsp                             = 5845
	FishBattleEndReq                               = 5841
	FishBattleEndRsp                               = 5842
	FishBiteReq                                    = 5844
	FishBiteRsp                                    = 5849
	FishCastRodReq                                 = 5802
	FishCastRodRsp                                 = 5831
	FishChosenNotify                               = 5829
	FishEscapeNotify                               = 5822
	FishPoolDataNotify                             = 5848
	FishingGallerySettleNotify                     = 8780
	FleurFairBalloonSettleNotify                   = 2099
	FleurFairBuffEnergyNotify                      = 5324
	FleurFairFallSettleNotify                      = 2017
	FleurFairFinishGalleryStageNotify              = 5342
	FleurFairMusicGameSettleReq                    = 2194
	FleurFairMusicGameSettleRsp                    = 2113
	FleurFairMusicGameStartReq                     = 2167
	FleurFairMusicGameStartRsp                     = 2079
	FleurFairReplayMiniGameReq                     = 2181
	FleurFairReplayMiniGameRsp                     = 2052
	FleurFairStageSettleNotify                     = 5356
	FlightActivityRestartReq                       = 2037
	FlightActivityRestartRsp                       = 2165
	FlightActivitySettleNotify                     = 2195
	FocusAvatarReq                                 = 1654
	FocusAvatarRsp                                 = 1681
	ForceAddPlayerFriendReq                        = 4057
	ForceAddPlayerFriendRsp                        = 4100
	ForceDragAvatarNotify                          = 3235
	ForceDragBackTransferNotify                    = 3145
	ForgeDataNotify                                = 680
	ForgeFormulaDataNotify                         = 689
	ForgeGetQueueDataReq                           = 646
	ForgeGetQueueDataRsp                           = 641
	ForgeQueueDataNotify                           = 676
	ForgeQueueManipulateReq                        = 624
	ForgeQueueManipulateRsp                        = 656
	ForgeStartReq                                  = 649
	ForgeStartRsp                                  = 691
	FoundationNotify                               = 847
	FoundationReq                                  = 805
	FoundationRsp                                  = 882
	FriendInfoChangeNotify                         = 4032
	FungusCaptureSettleNotify                      = 5506
	FungusCultivateReq                             = 21749
	FungusCultivateRsp                             = 23532
	FungusFighterClearTrainingRuntimeDataReq       = 24137
	FungusFighterClearTrainingRuntimeDataRsp       = 22991
	FungusFighterPlotInfoNotify                    = 22174
	FungusFighterRestartTrainingDungeonReq         = 23980
	FungusFighterRestartTrainingDungeonRsp         = 22890
	FungusFighterRuntimeDataNotify                 = 24674
	FungusFighterTrainingGallerySettleNotify       = 23931
	FungusFighterTrainingInfoNotify                = 5595
	FungusFighterTrainingSelectFungusReq           = 23903
	FungusFighterTrainingSelectFungusRsp           = 21570
	FungusFighterUseBackupFungusReq                = 21266
	FungusFighterUseBackupFungusRsp                = 23428
	FungusRenameReq                                = 22006
	FungusRenameRsp                                = 20066
	FunitureMakeInfoChangeNotify                   = 4898
	FurnitureCurModuleArrangeCountNotify           = 4498
	FurnitureMakeBeHelpedNotify                    = 4578
	FurnitureMakeCancelReq                         = 4555
	FurnitureMakeCancelRsp                         = 4683
	FurnitureMakeFinishNotify                      = 4841
	FurnitureMakeHelpReq                           = 4865
	FurnitureMakeHelpRsp                           = 4756
	FurnitureMakeReq                               = 4477
	FurnitureMakeRsp                               = 4782
	FurnitureMakeStartReq                          = 4633
	FurnitureMakeStartRsp                          = 4729
	GCGApplyInviteBattleNotify                     = 7820
	GCGApplyInviteBattleReq                        = 7730
	GCGApplyInviteBattleRsp                        = 7304
	GCGAskDuelReq                                  = 7237
	GCGAskDuelRsp                                  = 7869
	GCGBasicDataNotify                             = 7319
	GCGBossChallengeUpdateNotify                   = 7073
	GCGChallengeUpdateNotify                       = 7268
	GCGClientSettleReq                             = 7506
	GCGClientSettleRsp                             = 7105
	GCGDSCardBackUnlockNotify                      = 7265
	GCGDSCardFaceUnlockNotify                      = 7049
	GCGDSCardNumChangeNotify                       = 7358
	GCGDSCardProficiencyNotify                     = 7680
	GCGDSChangeCardBackReq                         = 7292
	GCGDSChangeCardBackRsp                         = 7044
	GCGDSChangeCardFaceReq                         = 7169
	GCGDSChangeCardFaceRsp                         = 7331
	GCGDSChangeCurDeckReq                          = 7131
	GCGDSChangeCurDeckRsp                          = 7301
	GCGDSChangeDeckNameReq                         = 7432
	GCGDSChangeDeckNameRsp                         = 7916
	GCGDSChangeFieldReq                            = 7541
	GCGDSChangeFieldRsp                            = 7444
	GCGDSCurDeckChangeNotify                       = 7796
	GCGDSDataNotify                                = 7122
	GCGDSDeckSaveReq                               = 7104
	GCGDSDeckSaveRsp                               = 7269
	GCGDSDeckUnlockNotify                          = 7732
	GCGDSDeleteDeckReq                             = 7988
	GCGDSDeleteDeckRsp                             = 7524
	GCGDSFieldUnlockNotify                         = 7333
	GCGGameBriefDataNotify                         = 7539
	GCGGrowthLevelNotify                           = 7736
	GCGGrowthLevelRewardNotify                     = 7477
	GCGGrowthLevelTakeRewardReq                    = 7051
	GCGGrowthLevelTakeRewardRsp                    = 7670
	GCGHeartBeatNotify                             = 7224
	GCGInitFinishReq                               = 7684
	GCGInitFinishRsp                               = 7433
	GCGInviteBattleNotify                          = 7692
	GCGInviteGuestBattleReq                        = 7783
	GCGInviteGuestBattleRsp                        = 7251
	GCGLevelChallengeFinishNotify                  = 7629
	GCGLevelChallengeNotify                        = 7055
	GCGMessagePackNotify                           = 7516
	GCGNewCardInfoNotify                           = 7203
	GCGOperationReq                                = 7107
	GCGOperationRsp                                = 7600
	GCGResourceStateNotify                         = 7876
	GCGSettleNotify                                = 7769
	GCGSettleOptionReq                             = 7124
	GCGSettleOptionRsp                             = 7735
	GCGSkillPreviewAskReq                          = 7509
	GCGSkillPreviewAskRsp                          = 7409
	GCGSkillPreviewNotify                          = 7503
	GCGStartChallengeReq                           = 7595
	GCGStartChallengeRsp                           = 7763
	GCGTCInviteReq                                 = 7922
	GCGTCInviteRsp                                 = 7328
	GCGTCTavernChallengeDataNotify                 = 7294
	GCGTCTavernChallengeUpdateNotify               = 7184
	GCGTCTavernInfoNotify                          = 7011
	GCGTavernNpcInfoNotify                         = 7290
	GCGWeekChallengeInfoNotify                     = 7615
	GCGWorldChallengeUnlockNotify                  = 7204
	GMShowNavMeshReq                               = 2357
	GMShowNavMeshRsp                               = 2400
	GMShowObstacleReq                              = 2361
	GMShowObstacleRsp                              = 2329
	GachaActivityCreateRobotReq                    = 8614
	GachaActivityCreateRobotRsp                    = 8610
	GachaActivityNextStageReq                      = 8257
	GachaActivityNextStageRsp                      = 8918
	GachaActivityPercentNotify                     = 8450
	GachaActivityResetReq                          = 8163
	GachaActivityResetRsp                          = 8240
	GachaActivityTakeRewardReq                     = 8930
	GachaActivityTakeRewardRsp                     = 8768
	GachaActivityUpdateElemNotify                  = 8919
	GachaOpenWishNotify                            = 1503
	GachaSimpleInfoNotify                          = 1590
	GachaWishReq                                   = 1507
	GachaWishRsp                                   = 1521
	GadgetAutoPickDropInfoNotify                   = 897
	GadgetChainLevelChangeNotify                   = 822
	GadgetChainLevelUpdateNotify                   = 853
	GadgetChangeLevelTagReq                        = 843
	GadgetChangeLevelTagRsp                        = 874
	GadgetCustomTreeInfoNotify                     = 850
	GadgetGeneralRewardInfoNotify                  = 848
	GadgetInteractReq                              = 872
	GadgetInteractRsp                              = 898
	GadgetPlayDataNotify                           = 831
	GadgetPlayStartNotify                          = 873
	GadgetPlayStopNotify                           = 899
	GadgetPlayUidOpNotify                          = 875
	GadgetStateNotify                              = 812
	GadgetTalkChangeNotify                         = 839
	GalleryBalloonScoreNotify                      = 5512
	GalleryBalloonShootNotify                      = 5598
	GalleryBounceConjuringHitNotify                = 5505
	GalleryBrokenFloorFallNotify                   = 5575
	GalleryBulletHitNotify                         = 5531
	GalleryCrystalLinkBuffInfoNotify               = 5539
	GalleryCrystalLinkKillMonsterNotify            = 5547
	GalleryFallCatchNotify                         = 5507
	GalleryFallScoreNotify                         = 5521
	GalleryFlowerCatchNotify                       = 5573
	GalleryIslandPartyDownHillInfoNotify           = 5522
	GalleryPreStartNotify                          = 5599
	GalleryStartNotify                             = 5572
	GalleryStopNotify                              = 5535
	GallerySumoKillMonsterNotify                   = 5582
	GalleryWillStartCountdownNotify                = 5594
	GearActivityFinishPlayGearReq                  = 21834
	GearActivityFinishPlayGearRsp                  = 21800
	GearActivityFinishPlayPictureReq               = 21054
	GearActivityFinishPlayPictureRsp               = 21851
	GearActivityStartPlayGearReq                   = 23467
	GearActivityStartPlayGearRsp                   = 21025
	GearActivityStartPlayPictureReq                = 24550
	GearActivityStartPlayPictureRsp                = 23388
	GetActivityInfoReq                             = 2095
	GetActivityInfoRsp                             = 2041
	GetActivityScheduleReq                         = 2136
	GetActivityScheduleRsp                         = 2107
	GetActivityShopSheetInfoReq                    = 703
	GetActivityShopSheetInfoRsp                    = 790
	GetAllActivatedBargainDataReq                  = 463
	GetAllActivatedBargainDataRsp                  = 495
	GetAllH5ActivityInfoReq                        = 5668
	GetAllH5ActivityInfoRsp                        = 5676
	GetAllMailNotify                               = 1497
	GetAllMailReq                                  = 1431
	GetAllMailResultNotify                         = 1481
	GetAllMailRsp                                  = 1475
	GetAllSceneGalleryInfoReq                      = 5503
	GetAllSceneGalleryInfoRsp                      = 5590
	GetAllUnlockNameCardReq                        = 4027
	GetAllUnlockNameCardRsp                        = 4094
	GetAreaExplorePointReq                         = 241
	GetAreaExplorePointRsp                         = 249
	GetAuthSalesmanInfoReq                         = 2070
	GetAuthSalesmanInfoRsp                         = 2004
	GetAuthkeyReq                                  = 1490
	GetAuthkeyRsp                                  = 1473
	GetBargainDataReq                              = 488
	GetBargainDataRsp                              = 426
	GetBattlePassProductReq                        = 2644
	GetBattlePassProductRsp                        = 2649
	GetBlossomBriefInfoListReq                     = 2772
	GetBlossomBriefInfoListRsp                     = 2798
	GetBonusActivityRewardReq                      = 2581
	GetBonusActivityRewardRsp                      = 2505
	GetChatEmojiCollectionReq                      = 4068
	GetChatEmojiCollectionRsp                      = 4033
	GetCityHuntingOfferReq                         = 4325
	GetCityHuntingOfferRsp                         = 4307
	GetCityReputationInfoReq                       = 2872
	GetCityReputationInfoRsp                       = 2898
	GetCityReputationMapInfoReq                    = 2875
	GetCityReputationMapInfoRsp                    = 2848
	GetCompoundDataReq                             = 141
	GetCompoundDataRsp                             = 149
	GetCustomDungeonReq                            = 6209
	GetCustomDungeonRsp                            = 6227
	GetDailyDungeonEntryInfoReq                    = 930
	GetDailyDungeonEntryInfoRsp                    = 967
	GetDungeonEntryExploreConditionReq             = 3165
	GetDungeonEntryExploreConditionRsp             = 3269
	GetExpeditionAssistInfoListReq                 = 2150
	GetExpeditionAssistInfoListRsp                 = 2035
	GetFriendShowAvatarInfoReq                     = 4070
	GetFriendShowAvatarInfoRsp                     = 4017
	GetFriendShowNameCardInfoReq                   = 4061
	GetFriendShowNameCardInfoRsp                   = 4029
	GetFurnitureCurModuleArrangeCountReq           = 4711
	GetGachaInfoReq                                = 1572
	GetGachaInfoRsp                                = 1598
	GetGameplayRecommendationReq                   = 151
	GetGameplayRecommendationRsp                   = 123
	GetHomeExchangeWoodInfoReq                     = 4473
	GetHomeExchangeWoodInfoRsp                     = 4659
	GetHomeLevelUpRewardReq                        = 4557
	GetHomeLevelUpRewardRsp                        = 4603
	GetHuntingOfferRewardReq                       = 4302
	GetHuntingOfferRewardRsp                       = 4331
	GetInvestigationMonsterReq                     = 1901
	GetInvestigationMonsterRsp                     = 1910
	GetMailItemReq                                 = 1435
	GetMailItemRsp                                 = 1407
	GetMapAreaReq                                  = 3108
	GetMapAreaRsp                                  = 3328
	GetMapMarkTipsReq                              = 3463
	GetMapMarkTipsRsp                              = 3327
	GetMechanicusInfoReq                           = 3972
	GetMechanicusInfoRsp                           = 3998
	GetNextResourceInfoReq                         = 192
	GetNextResourceInfoRsp                         = 120
	GetOnlinePlayerInfoReq                         = 82
	GetOnlinePlayerInfoRsp                         = 47
	GetOnlinePlayerListReq                         = 90
	GetOnlinePlayerListRsp                         = 73
	GetOpActivityInfoReq                           = 5172
	GetOpActivityInfoRsp                           = 5198
	GetParentQuestVideoKeyReq                      = 470
	GetParentQuestVideoKeyRsp                      = 417
	GetPlayerAskFriendListReq                      = 4018
	GetPlayerAskFriendListRsp                      = 4066
	GetPlayerBlacklistReq                          = 4049
	GetPlayerBlacklistRsp                          = 4091
	GetPlayerFriendListReq                         = 4072
	GetPlayerFriendListRsp                         = 4098
	GetPlayerHomeCompInfoReq                       = 4597
	GetPlayerMpModeAvailabilityReq                 = 1844
	GetPlayerMpModeAvailabilityRsp                 = 1849
	GetPlayerSocialDetailReq                       = 4073
	GetPlayerSocialDetailRsp                       = 4099
	GetPlayerTokenReq                              = 172
	GetPlayerTokenRsp                              = 198
	GetPushTipsRewardReq                           = 2227
	GetPushTipsRewardRsp                           = 2294
	GetQuestLackingResourceReq                     = 467
	GetQuestLackingResourceRsp                     = 458
	GetQuestTalkHistoryReq                         = 490
	GetQuestTalkHistoryRsp                         = 473
	GetRecentMpPlayerListReq                       = 4034
	GetRecentMpPlayerListRsp                       = 4050
	GetRecommendCustomDungeonReq                   = 6235
	GetRecommendCustomDungeonRsp                   = 6248
	GetRegionSearchReq                             = 5602
	GetReunionMissionInfoReq                       = 5094
	GetReunionMissionInfoRsp                       = 5099
	GetReunionPrivilegeInfoReq                     = 5097
	GetReunionPrivilegeInfoRsp                     = 5087
	GetReunionSignInInfoReq                        = 5052
	GetReunionSignInInfoRsp                        = 5081
	GetRogueDairyRepairInfoReq                     = 8014
	GetRogueDairyRepairInfoRsp                     = 8447
	GetSceneAreaReq                                = 265
	GetSceneAreaRsp                                = 204
	GetSceneNpcPositionReq                         = 535
	GetSceneNpcPositionRsp                         = 507
	GetScenePerformanceReq                         = 3419
	GetScenePerformanceRsp                         = 3137
	GetScenePointReq                               = 297
	GetScenePointRsp                               = 281
	GetShopReq                                     = 772
	GetShopRsp                                     = 798
	GetShopmallDataReq                             = 707
	GetShopmallDataRsp                             = 721
	GetSignInRewardReq                             = 2507
	GetSignInRewardRsp                             = 2521
	GetStoreCustomDungeonReq                       = 6250
	GetStoreCustomDungeonRsp                       = 6212
	GetUgcBriefInfoReq                             = 6325
	GetUgcBriefInfoRsp                             = 6307
	GetUgcReq                                      = 6326
	GetUgcRsp                                      = 6318
	GetWidgetSlotReq                               = 4253
	GetWidgetSlotRsp                               = 4254
	GetWorldMpInfoReq                              = 3391
	GetWorldMpInfoRsp                              = 3320
	GiveUpRoguelikeDungeonCardReq                  = 8353
	GiveUpRoguelikeDungeonCardRsp                  = 8497
	GivingRecordChangeNotify                       = 187
	GivingRecordNotify                             = 116
	GlobalBuildingInfoNotify                       = 5320
	GmTalkNotify                                   = 94
	GmTalkReq                                      = 98
	GmTalkRsp                                      = 12
	GrantRewardNotify                              = 663
	GravenInnocenceEditCarveCombinationReq         = 23107
	GravenInnocenceEditCarveCombinationRsp         = 20702
	GravenInnocencePhotoFinishReq                  = 21750
	GravenInnocencePhotoFinishRsp                  = 23948
	GravenInnocencePhotoReminderNotify             = 23864
	GravenInnocenceRaceRestartReq                  = 22882
	GravenInnocenceRaceRestartRsp                  = 21880
	GravenInnocenceRaceSettleNotify                = 20681
	GroupLinkAllNotify                             = 5776
	GroupLinkChangeNotify                          = 5768
	GroupLinkDeleteNotify                          = 5775
	GroupLinkMarkUpdateNotify                      = 5757
	GroupSuiteNotify                               = 3257
	GroupUnloadNotify                              = 3344
	GuestBeginEnterSceneNotify                     = 3031
	GuestPostEnterSceneNotify                      = 3144
	H5ActivityIdsNotify                            = 5675
	HideAndSeekPlayerReadyNotify                   = 5302
	HideAndSeekPlayerSetAvatarNotify               = 5319
	HideAndSeekSelectAvatarReq                     = 5330
	HideAndSeekSelectAvatarRsp                     = 5367
	HideAndSeekSelectSkillReq                      = 8183
	HideAndSeekSelectSkillRsp                      = 8088
	HideAndSeekSetReadyReq                         = 5358
	HideAndSeekSetReadyRsp                         = 5370
	HideAndSeekSettleNotify                        = 5317
	HitClientTrivialNotify                         = 244
	HitTreeNotify                                  = 3019
	HomeAllUnlockedBgmIdListNotify                 = 4608
	HomeAvatarAllFinishRewardNotify                = 4741
	HomeAvatarCostumeChangeNotify                  = 4748
	HomeAvatarRewardEventGetReq                    = 4551
	HomeAvatarRewardEventGetRsp                    = 4833
	HomeAvatarRewardEventNotify                    = 4852
	HomeAvatarSummonAllEventNotify                 = 4808
	HomeAvatarSummonEventReq                       = 4806
	HomeAvatarSummonEventRsp                       = 4817
	HomeAvatarSummonFinishReq                      = 4629
	HomeAvatarSummonFinishRsp                      = 4696
	HomeAvatarTalkFinishInfoNotify                 = 4896
	HomeAvatarTalkReq                              = 4688
	HomeAvatarTalkRsp                              = 4464
	HomeAvtarAllFinishRewardNotify                 = 4453
	HomeBalloonGalleryScoreNotify                  = 4654
	HomeBalloonGallerySettleNotify                 = 4811
	HomeBasicInfoNotify                            = 4885
	HomeBlockNotify                                = 4543
	HomeBlueprintInfoNotify                        = 4765
	HomeChangeBgmNotify                            = 4872
	HomeChangeBgmReq                               = 4558
	HomeChangeBgmRsp                               = 4488
	HomeChangeEditModeReq                          = 4564
	HomeChangeEditModeRsp                          = 4559
	HomeChangeModuleReq                            = 4809
	HomeChangeModuleRsp                            = 4596
	HomeChooseModuleReq                            = 4524
	HomeChooseModuleRsp                            = 4648
	HomeClearGroupRecordReq                        = 4759
	HomeClearGroupRecordRsp                        = 4605
	HomeComfortInfoNotify                          = 4699
	HomeCreateBlueprintReq                         = 4619
	HomeCreateBlueprintRsp                         = 4606
	HomeCustomFurnitureInfoNotify                  = 4712
	HomeDeleteBlueprintReq                         = 4502
	HomeDeleteBlueprintRsp                         = 4586
	HomeEditCustomFurnitureReq                     = 4724
	HomeEditCustomFurnitureRsp                     = 4496
	HomeEnterEditModeFinishReq                     = 4537
	HomeEnterEditModeFinishRsp                     = 4615
	HomeExchangeWoodReq                            = 4576
	HomeExchangeWoodRsp                            = 4622
	HomeFishFarmingInfoNotify                      = 4677
	HomeGalleryInPlayingNotify                     = 5553
	HomeGetArrangementInfoReq                      = 4848
	HomeGetArrangementInfoRsp                      = 4844
	HomeGetBasicInfoReq                            = 4655
	HomeGetBlueprintSlotInfoReq                    = 4584
	HomeGetBlueprintSlotInfoRsp                    = 4662
	HomeGetFishFarmingInfoReq                      = 4476
	HomeGetFishFarmingInfoRsp                      = 4678
	HomeGetGroupRecordReq                          = 4523
	HomeGetGroupRecordRsp                          = 4538
	HomeGetOnlineStatusReq                         = 4820
	HomeGetOnlineStatusRsp                         = 4705
	HomeKickPlayerReq                              = 4870
	HomeKickPlayerRsp                              = 4691
	HomeLimitedShopBuyGoodsReq                     = 4760
	HomeLimitedShopBuyGoodsRsp                     = 4750
	HomeLimitedShopGoodsListReq                    = 4552
	HomeLimitedShopGoodsListRsp                    = 4546
	HomeLimitedShopInfoChangeNotify                = 4790
	HomeLimitedShopInfoNotify                      = 4887
	HomeLimitedShopInfoReq                         = 4825
	HomeLimitedShopInfoRsp                         = 4796
	HomeMarkPointNotify                            = 4474
	HomeModuleSeenReq                              = 4499
	HomeModuleSeenRsp                              = 4821
	HomeModuleUnlockNotify                         = 4560
	HomeNewUnlockedBgmIdListNotify                 = 4847
	HomePictureFrameInfoNotify                     = 4878
	HomePlantFieldNotify                           = 4549
	HomePlantInfoNotify                            = 4587
	HomePlantInfoReq                               = 4647
	HomePlantInfoRsp                               = 4701
	HomePlantSeedReq                               = 4804
	HomePlantSeedRsp                               = 4556
	HomePlantWeedReq                               = 4640
	HomePlantWeedRsp                               = 4527
	HomePreChangeEditModeNotify                    = 4639
	HomePreviewBlueprintReq                        = 4478
	HomePreviewBlueprintRsp                        = 4738
	HomePriorCheckNotify                           = 4599
	HomeRacingGallerySettleNotify                  = 4805
	HomeResourceNotify                             = 4892
	HomeResourceTakeFetterExpReq                   = 4768
	HomeResourceTakeFetterExpRsp                   = 4645
	HomeResourceTakeHomeCoinReq                    = 4479
	HomeResourceTakeHomeCoinRsp                    = 4541
	HomeSaveArrangementNoChangeReq                 = 4704
	HomeSaveArrangementNoChangeRsp                 = 4668
	HomeSceneInitFinishReq                         = 4674
	HomeSceneInitFinishRsp                         = 4505
	HomeSceneJumpReq                               = 4528
	HomeSceneJumpRsp                               = 4698
	HomeScenePointFishFarmingInfoNotify            = 4547
	HomeSearchBlueprintReq                         = 4889
	HomeSearchBlueprintRsp                         = 4593
	HomeSeekFurnitureGalleryScoreNotify            = 4583
	HomeSetBlueprintFriendOptionReq                = 4554
	HomeSetBlueprintFriendOptionRsp                = 4604
	HomeSetBlueprintSlotOptionReq                  = 4798
	HomeSetBlueprintSlotOptionRsp                  = 4786
	HomeTransferReq                                = 4726
	HomeTransferRsp                                = 4616
	HomeUpdateArrangementInfoReq                   = 4510
	HomeUpdateArrangementInfoRsp                   = 4757
	HomeUpdateFishFarmingInfoReq                   = 4544
	HomeUpdateFishFarmingInfoRsp                   = 4857
	HomeUpdatePictureFrameInfoReq                  = 4486
	HomeUpdatePictureFrameInfoRsp                  = 4641
	HomeUpdateScenePointFishFarmingInfoReq         = 4511
	HomeUpdateScenePointFishFarmingInfoRsp         = 4540
	HostPlayerNotify                               = 312
	HuntingFailNotify                              = 4320
	HuntingGiveUpReq                               = 4341
	HuntingGiveUpRsp                               = 4342
	HuntingOngoingNotify                           = 4345
	HuntingRevealClueNotify                        = 4322
	HuntingRevealFinalNotify                       = 4344
	HuntingStartNotify                             = 4329
	HuntingSuccessNotify                           = 4349
	InBattleMechanicusBuildingPointsNotify         = 5303
	InBattleMechanicusCardResultNotify             = 5397
	InBattleMechanicusConfirmCardNotify            = 5348
	InBattleMechanicusConfirmCardReq               = 5331
	InBattleMechanicusConfirmCardRsp               = 5375
	InBattleMechanicusEscapeMonsterNotify          = 5307
	InBattleMechanicusLeftMonsterNotify            = 5321
	InBattleMechanicusPickCardNotify               = 5399
	InBattleMechanicusPickCardReq                  = 5390
	InBattleMechanicusPickCardRsp                  = 5373
	InBattleMechanicusSettleNotify                 = 5305
	InstableSprayEnterDungeonReq                   = 24312
	InstableSprayEnterDungeonRsp                   = 23381
	InstableSprayGalleryInfoNotify                 = 5588
	InstableSprayLevelFinishNotify                 = 21961
	InstableSprayRestartDungeonReq                 = 23678
	InstableSprayRestartDungeonRsp                 = 24923
	InstableSpraySwitchTeamReq                     = 24857
	InstableSpraySwitchTeamRsp                     = 24152
	InteractDailyDungeonInfoNotify                 = 919
	InterpretInferenceWordReq                      = 419
	InterpretInferenceWordRsp                      = 461
	InterruptGalleryReq                            = 5548
	InterruptGalleryRsp                            = 5597
	InvestigationMonsterUpdateNotify               = 1906
	InvestigationQuestDailyNotify                  = 1921
	InvestigationReadQuestDailyNotify              = 1902
	IrodoriChessEquipCardReq                       = 8561
	IrodoriChessEquipCardRsp                       = 8308
	IrodoriChessLeftMonsterNotify                  = 5338
	IrodoriChessPlayerInfoNotify                   = 5364
	IrodoriChessUnequipCardReq                     = 8057
	IrodoriChessUnequipCardRsp                     = 8817
	IrodoriEditFlowerCombinationReq                = 8608
	IrodoriEditFlowerCombinationRsp                = 8833
	IrodoriFillPoetryReq                           = 8129
	IrodoriFillPoetryRsp                           = 8880
	IrodoriMasterGalleryCgEndNotify                = 8061
	IrodoriMasterGallerySettleNotify               = 8340
	IrodoriMasterStartGalleryReq                   = 8165
	IrodoriMasterStartGalleryRsp                   = 8381
	IrodoriScanEntityReq                           = 8767
	IrodoriScanEntityRsp                           = 8026
	IslandPartyRaftInfoNotify                      = 5565
	IslandPartySailInfoNotify                      = 5504
	IslandPartySettleNotify                        = 24601
	ItemAddHintNotify                              = 607
	ItemCdGroupTimeNotify                          = 634
	ItemGivingReq                                  = 140
	ItemGivingRsp                                  = 118
	JoinHomeWorldFailNotify                        = 4530
	JoinPlayerFailNotify                           = 236
	JoinPlayerSceneReq                             = 292
	JoinPlayerSceneRsp                             = 220
	KeepAliveNotify                                = 72
	LanternRiteDoFireworksReformReq                = 8226
	LanternRiteDoFireworksReformRsp                = 8657
	LanternRiteEndFireworksReformReq               = 8277
	LanternRiteEndFireworksReformRsp               = 8933
	LanternRiteStartFireworksReformReq             = 8518
	LanternRiteStartFireworksReformRsp             = 8862
	LanternRiteTakeSkinRewardReq                   = 8826
	LanternRiteTakeSkinRewardRsp                   = 8777
	LastPacketPrintNotify                          = 88
	LaunchFireworksReq                             = 6090
	LaunchFireworksRsp                             = 6057
	LeaveSceneReq                                  = 298
	LeaveSceneRsp                                  = 212
	LeaveWorldNotify                               = 3017
	LevelTagDataNotify                             = 3314
	LevelupCityReq                                 = 216
	LevelupCityRsp                                 = 287
	LifeStateChangeNotify                          = 1298
	LikeCustomDungeonReq                           = 6210
	LikeCustomDungeonRsp                           = 6219
	LiveEndNotify                                  = 806
	LiveStartNotify                                = 826
	LoadActivityTerrainNotify                      = 2029
	LuaEnvironmentEffectNotify                     = 3408
	LuaSetOptionNotify                             = 316
	LuminanceStoneChallengeSettleNotify            = 8186
	LunaRiteAreaFinishNotify                       = 8213
	LunaRiteGroupBundleRegisterNotify              = 8465
	LunaRiteHintPointRemoveNotify                  = 8787
	LunaRiteHintPointReq                           = 8195
	LunaRiteHintPointRsp                           = 8765
	LunaRiteSacrificeReq                           = 8805
	LunaRiteSacrificeRsp                           = 8080
	LunaRiteTakeSacrificeRewardReq                 = 8045
	LunaRiteTakeSacrificeRewardRsp                 = 8397
	MailChangeNotify                               = 1498
	MainCoopFailNotify                             = 1951
	MainCoopUpdateNotify                           = 1968
	MapAreaChangeNotify                            = 3378
	MarkEntityInMinMapNotify                       = 202
	MarkMapReq                                     = 3466
	MarkMapRsp                                     = 3079
	MarkNewNotify                                  = 1275
	MarkTargetInvestigationMonsterNotify           = 1915
	MassiveEntityElementOpBatchNotify              = 357
	MassiveEntityStateChangedNotify                = 370
	MaterialDeleteReturnNotify                     = 661
	MaterialDeleteUpdateNotify                     = 700
	McoinExchangeHcoinReq                          = 616
	McoinExchangeHcoinRsp                          = 687
	MechanicusCandidateTeamCreateReq               = 3981
	MechanicusCandidateTeamCreateRsp               = 3905
	MechanicusCloseNotify                          = 3921
	MechanicusCoinNotify                           = 3935
	MechanicusLevelupGearReq                       = 3973
	MechanicusLevelupGearRsp                       = 3999
	MechanicusOpenNotify                           = 3907
	MechanicusSequenceOpenNotify                   = 3912
	MechanicusUnlockGearReq                        = 3903
	MechanicusUnlockGearRsp                        = 3990
	MeetNpcReq                                     = 503
	MeetNpcRsp                                     = 590
	MetNpcIdListNotify                             = 521
	MichiaeMatsuriDarkPressureLevelUpdateNotify    = 8825
	MichiaeMatsuriGainCrystalExpUpdateNotify       = 8523
	MichiaeMatsuriInteractStatueReq                = 8718
	MichiaeMatsuriInteractStatueRsp                = 8449
	MichiaeMatsuriRemoveChallengeMarkNotify        = 8072
	MichiaeMatsuriRemoveChestMarkNotify            = 8726
	MichiaeMatsuriStartBossChallengeReq            = 8703
	MichiaeMatsuriStartBossChallengeRsp            = 8426
	MichiaeMatsuriStartDarkChallengeReq            = 8054
	MichiaeMatsuriStartDarkChallengeRsp            = 8791
	MichiaeMatsuriUnlockCrystalSkillReq            = 8345
	MichiaeMatsuriUnlockCrystalSkillRsp            = 8588
	MiracleRingDataNotify                          = 5225
	MiracleRingDeliverItemReq                      = 5229
	MiracleRingDeliverItemRsp                      = 5222
	MiracleRingDestroyNotify                       = 5244
	MiracleRingDropResultNotify                    = 5231
	MiracleRingTakeRewardReq                       = 5207
	MiracleRingTakeRewardRsp                       = 5202
	MistTrialDungeonFailNotify                     = 8135
	MistTrialFloorLevelNotify                      = 968
	MistTrialGetChallengeMissionReq                = 8893
	MistTrialGetChallengeMissionRsp                = 8508
	MistTrialGetDungeonExhibitionDataReq           = 8740
	MistTrialGetDungeonExhibitionDataRsp           = 8066
	MistTrialSelectAvatarAndEnterDungeonReq        = 8666
	MistTrialSelectAvatarAndEnterDungeonRsp        = 8239
	MistTrialSettleNotify                          = 8373
	MonsterAIConfigHashNotify                      = 3039
	MonsterAlertChangeNotify                       = 363
	MonsterForceAlertNotify                        = 395
	MonsterPointArrayRouteUpdateNotify             = 3410
	MonsterSummonTagNotify                         = 1372
	MpBlockNotify                                  = 1801
	MpPlayGuestReplyInviteReq                      = 1848
	MpPlayGuestReplyInviteRsp                      = 1850
	MpPlayGuestReplyNotify                         = 1812
	MpPlayInviteResultNotify                       = 1815
	MpPlayOwnerCheckReq                            = 1814
	MpPlayOwnerCheckRsp                            = 1847
	MpPlayOwnerInviteNotify                        = 1835
	MpPlayOwnerStartInviteReq                      = 1837
	MpPlayOwnerStartInviteRsp                      = 1823
	MpPlayPrepareInterruptNotify                   = 1813
	MpPlayPrepareNotify                            = 1833
	MultistagePlayEndNotify                        = 5355
	MultistagePlayFinishStageReq                   = 5398
	MultistagePlayFinishStageRsp                   = 5381
	MultistagePlayInfoNotify                       = 5372
	MultistagePlaySettleNotify                     = 5313
	MultistagePlayStageEndNotify                   = 5379
	MuqadasPotionActivityEnterDungeonReq           = 24602
	MuqadasPotionActivityEnterDungeonRsp           = 21804
	MuqadasPotionCaptureWeaknessReq                = 20011
	MuqadasPotionCaptureWeaknessRsp                = 24081
	MuqadasPotionDungeonSettleNotify               = 20005
	MuqadasPotionRestartDungeonReq                 = 22391
	MuqadasPotionRestartDungeonRsp                 = 21208
	MusicGameSettleReq                             = 8892
	MusicGameSettleRsp                             = 8673
	MusicGameStartReq                              = 8406
	MusicGameStartRsp                              = 8326
	NavMeshStatsNotify                             = 2316
	NicknameAuditConfigNotify                      = 152
	NightCrowGadgetObservationMatchReq             = 876
	NightCrowGadgetObservationMatchRsp             = 846
	NormalUidOpNotify                              = 5726
	NpcTalkReq                                     = 572
	NpcTalkRsp                                     = 598
	NpcTalkStateNotify                             = 430
	ObstacleModifyNotify                           = 2312
	OfferingInteractReq                            = 2918
	OfferingInteractRsp                            = 2908
	OneofGatherPointDetectorDataNotify             = 4297
	OpActivityDataNotify                           = 5112
	OpActivityStateNotify                          = 2572
	OpActivityUpdateNotify                         = 5135
	OpenBlossomCircleCampGuideNotify               = 2703
	OpenStateChangeNotify                          = 127
	OpenStateUpdateNotify                          = 193
	OrderDisplayNotify                             = 4131
	OrderFinishNotify                              = 4125
	OtherPlayerEnterHomeNotify                     = 4628
	OutStuckCustomDungeonReq                       = 6211
	OutStuckCustomDungeonRsp                       = 6234
	PSNBlackListNotify                             = 4040
	PSNFriendListNotify                            = 4087
	PSPlayerApplyEnterMpReq                        = 1841
	PSPlayerApplyEnterMpRsp                        = 1842
	ParentQuestInferenceDataNotify                 = 402
	PathfindingEnterSceneReq                       = 2307
	PathfindingEnterSceneRsp                       = 2321
	PathfindingPingNotify                          = 2335
	PersistentDungeonSwitchAvatarReq               = 1684
	PersistentDungeonSwitchAvatarRsp               = 1768
	PersonalLineAllDataReq                         = 474
	PersonalLineAllDataRsp                         = 476
	PersonalLineNewUnlockNotify                    = 442
	PersonalSceneJumpReq                           = 284
	PersonalSceneJumpRsp                           = 280
	PhotoActivityClientViewReq                     = 8709
	PhotoActivityClientViewRsp                     = 8983
	PhotoActivityFinishReq                         = 8921
	PhotoActivityFinishRsp                         = 8854
	PingReq                                        = 7
	PingRsp                                        = 21
	PlantFlowerAcceptAllGiveFlowerReq              = 8808
	PlantFlowerAcceptAllGiveFlowerRsp              = 8888
	PlantFlowerAcceptGiveFlowerReq                 = 8383
	PlantFlowerAcceptGiveFlowerRsp                 = 8567
	PlantFlowerEditFlowerCombinationReq            = 8843
	PlantFlowerEditFlowerCombinationRsp            = 8788
	PlantFlowerGetCanGiveFriendFlowerReq           = 8716
	PlantFlowerGetCanGiveFriendFlowerRsp           = 8766
	PlantFlowerGetFriendFlowerWishListReq          = 8126
	PlantFlowerGetFriendFlowerWishListRsp          = 8511
	PlantFlowerGetRecvFlowerListReq                = 8270
	PlantFlowerGetRecvFlowerListRsp                = 8374
	PlantFlowerGetSeedInfoReq                      = 8560
	PlantFlowerGetSeedInfoRsp                      = 8764
	PlantFlowerGiveFriendFlowerReq                 = 8846
	PlantFlowerGiveFriendFlowerRsp                 = 8386
	PlantFlowerHaveRecvFlowerNotify                = 8078
	PlantFlowerSetFlowerWishReq                    = 8547
	PlantFlowerSetFlowerWishRsp                    = 8910
	PlantFlowerTakeSeedRewardReq                   = 8968
	PlantFlowerTakeSeedRewardRsp                   = 8860
	PlatformChangeRouteNotify                      = 268
	PlatformStartRouteNotify                       = 218
	PlatformStopRouteNotify                        = 266
	PlayerAllowEnterMpAfterAgreeMatchNotify        = 4199
	PlayerApplyEnterHomeNotify                     = 4533
	PlayerApplyEnterHomeResultNotify               = 4468
	PlayerApplyEnterHomeResultReq                  = 4693
	PlayerApplyEnterHomeResultRsp                  = 4706
	PlayerApplyEnterMpAfterMatchAgreedNotify       = 4195
	PlayerApplyEnterMpNotify                       = 1826
	PlayerApplyEnterMpReq                          = 1818
	PlayerApplyEnterMpResultNotify                 = 1807
	PlayerApplyEnterMpResultReq                    = 1802
	PlayerApplyEnterMpResultRsp                    = 1831
	PlayerApplyEnterMpRsp                          = 1825
	PlayerCancelMatchReq                           = 4157
	PlayerCancelMatchRsp                           = 4152
	PlayerChatCDNotify                             = 3367
	PlayerChatNotify                               = 3010
	PlayerChatReq                                  = 3185
	PlayerChatRsp                                  = 3228
	PlayerCompoundMaterialBoostReq                 = 185
	PlayerCompoundMaterialBoostRsp                 = 125
	PlayerCompoundMaterialReq                      = 150
	PlayerCompoundMaterialRsp                      = 143
	PlayerConfirmMatchReq                          = 4172
	PlayerConfirmMatchRsp                          = 4194
	PlayerCookArgsReq                              = 166
	PlayerCookArgsRsp                              = 168
	PlayerCookReq                                  = 194
	PlayerCookRsp                                  = 188
	PlayerDataNotify                               = 190
	PlayerDeathZoneNotify                          = 6275
	PlayerEnterDungeonReq                          = 912
	PlayerEnterDungeonRsp                          = 935
	PlayerEnterSceneInfoNotify                     = 214
	PlayerEnterSceneNotify                         = 272
	PlayerEyePointStateNotify                      = 3051
	PlayerFishingDataNotify                        = 5835
	PlayerForceExitReq                             = 189
	PlayerForceExitRsp                             = 159
	PlayerGCGMatchConfirmNotify                    = 4185
	PlayerGCGMatchDismissNotify                    = 4173
	PlayerGameTimeNotify                           = 131
	PlayerGeneralMatchConfirmNotify                = 4192
	PlayerGeneralMatchDismissNotify                = 4191
	PlayerGetForceQuitBanInfoReq                   = 4164
	PlayerGetForceQuitBanInfoRsp                   = 4197
	PlayerHomeCompInfoNotify                       = 4880
	PlayerInjectFixNotify                          = 132
	PlayerInvestigationAllInfoNotify               = 1928
	PlayerInvestigationNotify                      = 1911
	PlayerInvestigationTargetNotify                = 1929
	PlayerLevelRewardUpdateNotify                  = 200
	PlayerLoginReq                                 = 112
	PlayerLoginRsp                                 = 135
	PlayerLogoutNotify                             = 103
	PlayerLogoutReq                                = 107
	PlayerLogoutRsp                                = 121
	PlayerLuaShellNotify                           = 133
	PlayerMatchAgreedResultNotify                  = 4170
	PlayerMatchInfoNotify                          = 4175
	PlayerMatchStopNotify                          = 4181
	PlayerMatchSuccNotify                          = 4179
	PlayerNicknameAuditDataNotify                  = 108
	PlayerNicknameNotify                           = 109
	PlayerOfferingDataNotify                       = 2923
	PlayerOfferingReq                              = 2907
	PlayerOfferingRsp                              = 2917
	PlayerPreEnterMpNotify                         = 1822
	PlayerPropChangeNotify                         = 139
	PlayerPropChangeReasonNotify                   = 1299
	PlayerPropNotify                               = 175
	PlayerQuitDungeonReq                           = 907
	PlayerQuitDungeonRsp                           = 921
	PlayerQuitFromHomeNotify                       = 4656
	PlayerQuitFromMpNotify                         = 1829
	PlayerRandomCookReq                            = 126
	PlayerRandomCookRsp                            = 163
	PlayerRechargeDataNotify                       = 4102
	PlayerReportReq                                = 4024
	PlayerReportRsp                                = 4056
	PlayerRoutineDataNotify                        = 3526
	PlayerSetLanguageReq                           = 142
	PlayerSetLanguageRsp                           = 130
	PlayerSetOnlyMPWithPSPlayerReq                 = 1820
	PlayerSetOnlyMPWithPSPlayerRsp                 = 1845
	PlayerSetPauseReq                              = 124
	PlayerSetPauseRsp                              = 156
	PlayerSignatureAuditDataNotify                 = 4060
	PlayerSignatureNotify                          = 4014
	PlayerStartMatchReq                            = 4176
	PlayerStartMatchRsp                            = 4168
	PlayerStoreNotify                              = 672
	PlayerTimeNotify                               = 191
	PlayerWorldSceneInfoListNotify                 = 3129
	PostEnterSceneReq                              = 3312
	PostEnterSceneRsp                              = 3184
	PotionEnterDungeonNotify                       = 8531
	PotionEnterDungeonReq                          = 8261
	PotionEnterDungeonRsp                          = 8482
	PotionResetChallengeReq                        = 8377
	PotionResetChallengeRsp                        = 8067
	PotionRestartDungeonReq                        = 8273
	PotionRestartDungeonRsp                        = 8062
	PotionSaveDungeonResultReq                     = 8192
	PotionSaveDungeonResultRsp                     = 8688
	PrivateChatNotify                              = 4962
	PrivateChatReq                                 = 5022
	PrivateChatRsp                                 = 5048
	ProfilePictureChangeNotify                     = 4016
	ProjectorOptionReq                             = 863
	ProjectorOptionRsp                             = 895
	ProudSkillChangeNotify                         = 1031
	ProudSkillExtraLevelNotify                     = 1081
	ProudSkillUpgradeReq                           = 1073
	ProudSkillUpgradeRsp                           = 1099
	PublishCustomDungeonReq                        = 6242
	PublishCustomDungeonRsp                        = 6214
	PublishUgcReq                                  = 6344
	PublishUgcRsp                                  = 6349
	PullPrivateChatReq                             = 4971
	PullPrivateChatRsp                             = 4953
	PullRecentChatReq                              = 5040
	PullRecentChatRsp                              = 5023
	PushTipsAllDataNotify                          = 2222
	PushTipsChangeNotify                           = 2265
	PushTipsReadFinishReq                          = 2204
	PushTipsReadFinishRsp                          = 2293
	QueryCodexMonsterBeKilledNumReq                = 4203
	QueryCodexMonsterBeKilledNumRsp                = 4209
	QueryPathReq                                   = 2372
	QueryPathRsp                                   = 2398
	QuestCreateEntityReq                           = 499
	QuestCreateEntityRsp                           = 431
	QuestDelNotify                                 = 412
	QuestDestroyEntityReq                          = 475
	QuestDestroyEntityRsp                          = 448
	QuestDestroyNpcReq                             = 422
	QuestDestroyNpcRsp                             = 465
	QuestGlobalVarNotify                           = 434
	QuestListNotify                                = 472
	QuestListUpdateNotify                          = 498
	QuestProgressUpdateNotify                      = 482
	QuestTransmitReq                               = 450
	QuestTransmitRsp                               = 443
	QuestUpdateQuestTimeVarNotify                  = 456
	QuestUpdateQuestVarNotify                      = 453
	QuestUpdateQuestVarReq                         = 447
	QuestUpdateQuestVarRsp                         = 439
	QuickOpenActivityReq                           = 8178
	QuickOpenActivityRsp                           = 8882
	QuickUseWidgetReq                              = 4299
	QuickUseWidgetRsp                              = 4270
	ReadMailNotify                                 = 1412
	ReadNicknameAuditReq                           = 177
	ReadNicknameAuditRsp                           = 137
	ReadPrivateChatReq                             = 5049
	ReadPrivateChatRsp                             = 4981
	ReadSignatureAuditReq                          = 4020
	ReadSignatureAuditRsp                          = 4064
	ReceivedTrialAvatarActivityRewardReq           = 2130
	ReceivedTrialAvatarActivityRewardRsp           = 2076
	RechargeReq                                    = 4126
	RechargeRsp                                    = 4118
	RedeemLegendaryKeyReq                          = 446
	RedeemLegendaryKeyRsp                          = 441
	ReformFireworksReq                             = 6036
	ReformFireworksRsp                             = 5929
	RefreshBackgroundAvatarReq                     = 1743
	RefreshBackgroundAvatarRsp                     = 1800
	RefreshEntityAuthNotify                        = 3259
	RefreshRogueDiaryCardReq                       = 8991
	RefreshRogueDiaryCardRsp                       = 8028
	RefreshRoguelikeDungeonCardReq                 = 8279
	RefreshRoguelikeDungeonCardRsp                 = 8349
	RegionSearchChangeRegionNotify                 = 5618
	RegionSearchNotify                             = 5626
	RegionalPlayInfoNotify                         = 6276
	ReliquaryDecomposeReq                          = 638
	ReliquaryDecomposeRsp                          = 611
	ReliquaryPromoteReq                            = 627
	ReliquaryPromoteRsp                            = 694
	ReliquaryUpgradeReq                            = 604
	ReliquaryUpgradeRsp                            = 693
	RemotePlayerWidgetNotify                       = 5995
	RemoveBlacklistReq                             = 4063
	RemoveBlacklistRsp                             = 4095
	RemoveCustomDungeonReq                         = 6249
	RemoveCustomDungeonRsp                         = 6220
	RemoveRandTaskInfoNotify                       = 161
	ReplayCustomDungeonReq                         = 6243
	ReplayCustomDungeonRsp                         = 6240
	ReportFightAntiCheatNotify                     = 368
	ReportTrackingIOInfoNotify                     = 4129
	RequestLiveInfoReq                             = 894
	RequestLiveInfoRsp                             = 888
	ReserveRogueDiaryAvatarReq                     = 8748
	ReserveRogueDiaryAvatarRsp                     = 8799
	ResetRogueDiaryPlayReq                         = 8127
	ResetRogueDiaryPlayRsp                         = 8948
	ResinCardDataUpdateNotify                      = 4149
	ResinChangeNotify                              = 642
	RestartEffigyChallengeReq                      = 2148
	RestartEffigyChallengeRsp                      = 2042
	ResumeRogueDiaryDungeonReq                     = 8838
	ResumeRogueDiaryDungeonRsp                     = 8989
	RetryCurRogueDiaryDungeonReq                   = 8398
	RetryCurRogueDiaryDungeonRsp                   = 8334
	ReunionActivateNotify                          = 5085
	ReunionBriefInfoReq                            = 5076
	ReunionBriefInfoRsp                            = 5068
	ReunionDailyRefreshNotify                      = 5100
	ReunionPrivilegeChangeNotify                   = 5098
	ReunionSettleNotify                            = 5073
	RobotPushPlayerDataNotify                      = 97
	RogueCellUpdateNotify                          = 8642
	RogueDiaryCoinAddNotify                        = 8602
	RogueDiaryDungeonInfoNotify                    = 8597
	RogueDiaryDungeonSettleNotify                  = 8895
	RogueDiaryRepairInfoNotify                     = 8641
	RogueDiaryReviveAvatarReq                      = 8038
	RogueDiaryReviveAvatarRsp                      = 8343
	RogueDiaryTiredAvatarNotify                    = 8514
	RogueDungeonPlayerCellChangeNotify             = 8347
	RogueFinishRepairReq                           = 8363
	RogueFinishRepairRsp                           = 8535
	RogueHealAvatarsReq                            = 8947
	RogueHealAvatarsRsp                            = 8949
	RogueResumeDungeonReq                          = 8795
	RogueResumeDungeonRsp                          = 8647
	RogueSwitchAvatarReq                           = 8201
	RogueSwitchAvatarRsp                           = 8915
	RoguelikeCardGachaNotify                       = 8925
	RoguelikeEffectDataNotify                      = 8222
	RoguelikeEffectViewReq                         = 8528
	RoguelikeEffectViewRsp                         = 8639
	RoguelikeGiveUpReq                             = 8660
	RoguelikeGiveUpRsp                             = 8139
	RoguelikeMistClearNotify                       = 8324
	RoguelikeRefreshCardCostUpdateNotify           = 8927
	RoguelikeResourceBonusPropUpdateNotify         = 8555
	RoguelikeRuneRecordUpdateNotify                = 8973
	RoguelikeSelectAvatarAndEnterDungeonReq        = 8457
	RoguelikeSelectAvatarAndEnterDungeonRsp        = 8538
	RoguelikeTakeStageFirstPassRewardReq           = 8421
	RoguelikeTakeStageFirstPassRewardRsp           = 8552
	SalesmanDeliverItemReq                         = 2138
	SalesmanDeliverItemRsp                         = 2104
	SalesmanTakeRewardReq                          = 2191
	SalesmanTakeRewardRsp                          = 2110
	SalesmanTakeSpecialRewardReq                   = 2145
	SalesmanTakeSpecialRewardRsp                   = 2124
	SalvageEscortRestartReq                        = 8396
	SalvageEscortRestartRsp                        = 8959
	SalvageEscortSettleNotify                      = 8499
	SalvagePreventRestartReq                       = 8367
	SalvagePreventRestartRsp                       = 8938
	SalvagePreventSettleNotify                     = 8231
	SaveCoopDialogReq                              = 2000
	SaveCoopDialogRsp                              = 1962
	SaveCustomDungeonRoomReq                       = 6225
	SaveCustomDungeonRoomRsp                       = 6207
	SaveMainCoopReq                                = 1975
	SaveMainCoopRsp                                = 1957
	SaveUgcReq                                     = 6329
	SaveUgcRsp                                     = 6322
	SceneAreaUnlockNotify                          = 293
	SceneAreaWeatherNotify                         = 230
	SceneAudioNotify                               = 3166
	SceneAvatarStaminaStepReq                      = 299
	SceneAvatarStaminaStepRsp                      = 231
	SceneCreateEntityReq                           = 288
	SceneCreateEntityRsp                           = 226
	SceneDataNotify                                = 3203
	SceneDestroyEntityReq                          = 263
	SceneDestroyEntityRsp                          = 295
	SceneEntitiesMoveCombineNotify                 = 3387
	SceneEntitiesMovesReq                          = 279
	SceneEntitiesMovesRsp                          = 255
	SceneEntityAppearNotify                        = 221
	SceneEntityDisappearNotify                     = 203
	SceneEntityDrownReq                            = 227
	SceneEntityDrownRsp                            = 294
	SceneEntityMoveNotify                          = 275
	SceneEntityMoveReq                             = 290
	SceneEntityMoveRsp                             = 273
	SceneEntityUpdateNotify                        = 3412
	SceneForceLockNotify                           = 234
	SceneForceUnlockNotify                         = 206
	SceneGalleryInfoNotify                         = 5581
	SceneGalleryVintageHuntingSettleNotify         = 20324
	SceneInitFinishReq                             = 235
	SceneInitFinishRsp                             = 207
	SceneKickPlayerNotify                          = 211
	SceneKickPlayerReq                             = 264
	SceneKickPlayerRsp                             = 238
	ScenePlayBattleInfoListNotify                  = 4431
	ScenePlayBattleInfoNotify                      = 4422
	ScenePlayBattleInterruptNotify                 = 4425
	ScenePlayBattleResultNotify                    = 4398
	ScenePlayBattleUidOpNotify                     = 4447
	ScenePlayGuestReplyInviteReq                   = 4353
	ScenePlayGuestReplyInviteRsp                   = 4440
	ScenePlayGuestReplyNotify                      = 4423
	ScenePlayInfoListNotify                        = 4381
	ScenePlayInviteResultNotify                    = 4449
	ScenePlayOutofRegionNotify                     = 4355
	ScenePlayOwnerCheckReq                         = 4448
	ScenePlayOwnerCheckRsp                         = 4362
	ScenePlayOwnerInviteNotify                     = 4371
	ScenePlayOwnerStartInviteReq                   = 4385
	ScenePlayOwnerStartInviteRsp                   = 4357
	ScenePlayerBackgroundAvatarRefreshNotify       = 3274
	ScenePlayerInfoNotify                          = 267
	ScenePlayerLocationNotify                      = 248
	ScenePlayerSoundNotify                         = 233
	ScenePointUnlockNotify                         = 247
	SceneRouteChangeNotify                         = 240
	SceneTeamUpdateNotify                          = 1775
	SceneTimeNotify                                = 245
	SceneTransToPointReq                           = 239
	SceneTransToPointRsp                           = 253
	SceneWeatherForecastReq                        = 3110
	SceneWeatherForecastRsp                        = 3012
	SeaLampCoinNotify                              = 2114
	SeaLampContributeItemReq                       = 2123
	SeaLampContributeItemRsp                       = 2139
	SeaLampFlyLampNotify                           = 2105
	SeaLampFlyLampReq                              = 2199
	SeaLampFlyLampRsp                              = 2192
	SeaLampPopularityNotify                        = 2032
	SeaLampTakeContributionRewardReq               = 2019
	SeaLampTakeContributionRewardRsp               = 2177
	SeaLampTakePhaseRewardReq                      = 2176
	SeaLampTakePhaseRewardRsp                      = 2190
	SealBattleBeginNotify                          = 289
	SealBattleEndNotify                            = 259
	SealBattleProgressNotify                       = 232
	SearchCustomDungeonReq                         = 6233
	SearchCustomDungeonRsp                         = 6215
	SeeMonsterReq                                  = 228
	SeeMonsterRsp                                  = 251
	SelectAsterMidDifficultyReq                    = 2134
	SelectAsterMidDifficultyRsp                    = 2180
	SelectEffigyChallengeConditionReq              = 2064
	SelectEffigyChallengeConditionRsp              = 2039
	SelectRoguelikeDungeonCardReq                  = 8085
	SelectRoguelikeDungeonCardRsp                  = 8138
	SelectWorktopOptionReq                         = 807
	SelectWorktopOptionRsp                         = 821
	ServerAnnounceNotify                           = 2197
	ServerAnnounceRevokeNotify                     = 2092
	ServerBuffChangeNotify                         = 361
	ServerCombatEndNotify                          = 1105
	ServerCondMeetQuestListUpdateNotify            = 406
	ServerDisconnectClientNotify                   = 184
	ServerGlobalValueChangeNotify                  = 1197
	ServerLogNotify                                = 31
	ServerMessageNotify                            = 5718
	ServerTimeNotify                               = 99
	ServerTryCancelGeneralMatchNotify              = 4187
	ServerUpdateGlobalValueNotify                  = 1148
	SetBattlePassViewedReq                         = 2641
	SetBattlePassViewedRsp                         = 2642
	SetChatEmojiCollectionReq                      = 4084
	SetChatEmojiCollectionRsp                      = 4080
	SetCodexPushtipsReadReq                        = 4208
	SetCodexPushtipsReadRsp                        = 4206
	SetCoopChapterViewedReq                        = 1965
	SetCoopChapterViewedRsp                        = 1963
	SetCurExpeditionChallengeIdReq                 = 2021
	SetCurExpeditionChallengeIdRsp                 = 2049
	SetEntityClientDataNotify                      = 3146
	SetEquipLockStateReq                           = 666
	SetEquipLockStateRsp                           = 668
	SetFriendEnterHomeOptionReq                    = 4494
	SetFriendEnterHomeOptionRsp                    = 4743
	SetFriendRemarkNameReq                         = 4042
	SetFriendRemarkNameRsp                         = 4030
	SetH5ActivityRedDotTimestampReq                = 5657
	SetH5ActivityRedDotTimestampRsp                = 5652
	SetIsAutoUnlockSpecificEquipReq                = 620
	SetIsAutoUnlockSpecificEquipRsp                = 664
	SetLimitOptimizationNotify                     = 8851
	SetNameCardReq                                 = 4004
	SetNameCardRsp                                 = 4093
	SetOpenStateReq                                = 165
	SetOpenStateRsp                                = 104
	SetPlayerBirthdayReq                           = 4048
	SetPlayerBirthdayRsp                           = 4097
	SetPlayerBornDataReq                           = 105
	SetPlayerBornDataRsp                           = 182
	SetPlayerHeadImageReq                          = 4082
	SetPlayerHeadImageRsp                          = 4047
	SetPlayerNameReq                               = 153
	SetPlayerNameRsp                               = 122
	SetPlayerPropReq                               = 197
	SetPlayerPropRsp                               = 181
	SetPlayerSignatureReq                          = 4081
	SetPlayerSignatureRsp                          = 4005
	SetSceneWeatherAreaReq                         = 254
	SetSceneWeatherAreaRsp                         = 283
	SetUpAvatarTeamReq                             = 1690
	SetUpAvatarTeamRsp                             = 1646
	SetUpLunchBoxWidgetReq                         = 4272
	SetUpLunchBoxWidgetRsp                         = 4294
	SetWidgetSlotReq                               = 4259
	SetWidgetSlotRsp                               = 4277
	ShowClientGuideNotify                          = 3005
	ShowClientTutorialNotify                       = 3305
	ShowCommonTipsNotify                           = 3352
	ShowMessageNotify                              = 35
	ShowTemplateReminderNotify                     = 3491
	SignInInfoReq                                  = 2512
	SignInInfoRsp                                  = 2535
	SignatureAuditConfigNotify                     = 4092
	SkyCrystalDetectorDataUpdateNotify             = 4287
	SocialDataNotify                               = 4043
	SpiceActivityFinishMakeSpiceReq                = 8096
	SpiceActivityFinishMakeSpiceRsp                = 8481
	SpiceActivityGivingRecordNotify                = 8407
	SpiceActivityProcessFoodReq                    = 8216
	SpiceActivityProcessFoodRsp                    = 8772
	SpringUseReq                                   = 1748
	SpringUseRsp                                   = 1642
	StartArenaChallengeLevelReq                    = 2127
	StartArenaChallengeLevelRsp                    = 2125
	StartBuoyantCombatGalleryReq                   = 8732
	StartBuoyantCombatGalleryRsp                   = 8680
	StartCoopPointReq                              = 1992
	StartCoopPointRsp                              = 1964
	StartEffigyChallengeReq                        = 2169
	StartEffigyChallengeRsp                        = 2173
	StartFishingReq                                = 5825
	StartFishingRsp                                = 5807
	StartRogueDiaryPlayReq                         = 8419
	StartRogueDiaryPlayRsp                         = 8385
	StartRogueDiaryRoomReq                         = 8159
	StartRogueDiaryRoomRsp                         = 8793
	StartRogueEliteCellChallengeReq                = 8242
	StartRogueEliteCellChallengeRsp                = 8958
	StartRogueNormalCellChallengeReq               = 8205
	StartRogueNormalCellChallengeRsp               = 8036
	StopReminderNotify                             = 3004
	StoreCustomDungeonReq                          = 6213
	StoreCustomDungeonRsp                          = 6201
	StoreItemChangeNotify                          = 612
	StoreItemDelNotify                             = 635
	StoreWeightLimitNotify                         = 698
	SubmitInferenceWordReq                         = 500
	SubmitInferenceWordRsp                         = 416
	SummerTimeFloatSignalPositionNotify            = 8077
	SummerTimeFloatSignalUpdateNotify              = 8781
	SummerTimeSprintBoatRestartReq                 = 8410
	SummerTimeSprintBoatRestartRsp                 = 8356
	SummerTimeSprintBoatSettleNotify               = 8651
	SummerTimeV2BoatSettleNotify                   = 8870
	SummerTimeV2RestartBoatGalleryReq              = 8476
	SummerTimeV2RestartBoatGalleryRsp              = 8004
	SummerTimeV2RestartDungeonReq                  = 8346
	SummerTimeV2RestartDungeonRsp                  = 8996
	SumoDungeonSettleNotify                        = 8291
	SumoEnterDungeonNotify                         = 8013
	SumoLeaveDungeonNotify                         = 8640
	SumoRestartDungeonReq                          = 8612
	SumoRestartDungeonRsp                          = 8214
	SumoSaveTeamReq                                = 8313
	SumoSaveTeamRsp                                = 8319
	SumoSelectTeamAndEnterDungeonReq               = 8215
	SumoSelectTeamAndEnterDungeonRsp               = 8193
	SumoSetNoSwitchPunishTimeNotify                = 8935
	SumoSwitchTeamReq                              = 8351
	SumoSwitchTeamRsp                              = 8525
	SyncScenePlayTeamEntityNotify                  = 3333
	SyncTeamEntityNotify                           = 317
	TakeAchievementGoalRewardReq                   = 2652
	TakeAchievementGoalRewardRsp                   = 2681
	TakeAchievementRewardReq                       = 2675
	TakeAchievementRewardRsp                       = 2657
	TakeAsterSpecialRewardReq                      = 2097
	TakeAsterSpecialRewardRsp                      = 2193
	TakeBackGivingItemReq                          = 171
	TakeBackGivingItemRsp                          = 145
	TakeBattlePassMissionPointReq                  = 2629
	TakeBattlePassMissionPointRsp                  = 2622
	TakeBattlePassRewardReq                        = 2602
	TakeBattlePassRewardRsp                        = 2631
	TakeCityReputationExploreRewardReq             = 2897
	TakeCityReputationExploreRewardRsp             = 2881
	TakeCityReputationLevelRewardReq               = 2812
	TakeCityReputationLevelRewardRsp               = 2835
	TakeCityReputationParentQuestReq               = 2821
	TakeCityReputationParentQuestRsp               = 2803
	TakeCompoundOutputReq                          = 174
	TakeCompoundOutputRsp                          = 176
	TakeCoopRewardReq                              = 1973
	TakeCoopRewardRsp                              = 1985
	TakeDeliveryDailyRewardReq                     = 2121
	TakeDeliveryDailyRewardRsp                     = 2162
	TakeEffigyFirstPassRewardReq                   = 2196
	TakeEffigyFirstPassRewardRsp                   = 2061
	TakeEffigyRewardReq                            = 2040
	TakeEffigyRewardRsp                            = 2007
	TakeFirstShareRewardReq                        = 4074
	TakeFirstShareRewardRsp                        = 4076
	TakeFurnitureMakeReq                           = 4772
	TakeFurnitureMakeRsp                           = 4769
	TakeHuntingOfferReq                            = 4326
	TakeHuntingOfferRsp                            = 4318
	TakeInvestigationRewardReq                     = 1912
	TakeInvestigationRewardRsp                     = 1922
	TakeInvestigationTargetRewardReq               = 1918
	TakeInvestigationTargetRewardRsp               = 1916
	TakeMaterialDeleteReturnReq                    = 629
	TakeMaterialDeleteReturnRsp                    = 657
	TakeOfferingLevelRewardReq                     = 2919
	TakeOfferingLevelRewardRsp                     = 2911
	TakePlayerLevelRewardReq                       = 129
	TakePlayerLevelRewardRsp                       = 157
	TakeRegionSearchRewardReq                      = 5625
	TakeRegionSearchRewardRsp                      = 5607
	TakeResinCardDailyRewardReq                    = 4122
	TakeResinCardDailyRewardRsp                    = 4144
	TakeReunionFirstGiftRewardReq                  = 5075
	TakeReunionFirstGiftRewardRsp                  = 5057
	TakeReunionMissionRewardReq                    = 5092
	TakeReunionMissionRewardRsp                    = 5064
	TakeReunionSignInRewardReq                     = 5079
	TakeReunionSignInRewardRsp                     = 5072
	TakeReunionWatcherRewardReq                    = 5070
	TakeReunionWatcherRewardRsp                    = 5095
	TakeoffEquipReq                                = 605
	TakeoffEquipRsp                                = 682
	TanukiTravelFinishGuideQuestNotify             = 8924
	TaskVarNotify                                  = 160
	TeamResonanceChangeNotify                      = 1082
	ToTheMoonAddObstacleReq                        = 6121
	ToTheMoonAddObstacleRsp                        = 6103
	ToTheMoonEnterSceneReq                         = 6135
	ToTheMoonEnterSceneRsp                         = 6107
	ToTheMoonObstaclesModifyNotify                 = 6199
	ToTheMoonPingNotify                            = 6112
	ToTheMoonQueryPathReq                          = 6172
	ToTheMoonQueryPathRsp                          = 6198
	ToTheMoonRemoveObstacleReq                     = 6190
	ToTheMoonRemoveObstacleRsp                     = 6173
	TowerAllDataReq                                = 2490
	TowerAllDataRsp                                = 2473
	TowerBriefDataNotify                           = 2472
	TowerBuffSelectReq                             = 2448
	TowerBuffSelectRsp                             = 2497
	TowerCurLevelRecordChangeNotify                = 2412
	TowerDailyRewardProgressChangeNotify           = 2435
	TowerEnterLevelReq                             = 2431
	TowerEnterLevelRsp                             = 2475
	TowerFloorRecordChangeNotify                   = 2498
	TowerGetFloorStarRewardReq                     = 2404
	TowerGetFloorStarRewardRsp                     = 2493
	TowerLevelEndNotify                            = 2495
	TowerLevelStarCondNotify                       = 2406
	TowerMiddleLevelChangeTeamNotify               = 2434
	TowerRecordHandbookReq                         = 2450
	TowerRecordHandbookRsp                         = 2443
	TowerSurrenderReq                              = 2422
	TowerSurrenderRsp                              = 2465
	TowerTeamSelectReq                             = 2421
	TowerTeamSelectRsp                             = 2403
	TreasureMapBonusChallengeNotify                = 2115
	TreasureMapCurrencyNotify                      = 2171
	TreasureMapDetectorDataNotify                  = 4300
	TreasureMapGuideTaskDoneNotify                 = 2119
	TreasureMapHostInfoNotify                      = 8681
	TreasureMapMpChallengeNotify                   = 2048
	TreasureMapPreTaskDoneNotify                   = 2152
	TreasureMapRegionActiveNotify                  = 2122
	TreasureMapRegionInfoNotify                    = 2185
	TreasureSeelieCollectOrbsNotify                = 20754
	TrialAvatarFirstPassDungeonNotify              = 2013
	TrialAvatarInDungeonIndexNotify                = 2186
	TriggerCreateGadgetToEquipPartNotify           = 350
	TriggerRoguelikeCurseNotify                    = 8412
	TriggerRoguelikeRuneReq                        = 8463
	TriggerRoguelikeRuneRsp                        = 8065
	TryCustomDungeonReq                            = 6245
	TryCustomDungeonRsp                            = 6241
	TryEnterHomeReq                                = 4482
	TryEnterHomeRsp                                = 4653
	TryEnterNextRogueDiaryDungeonReq               = 8280
	TryEnterNextRogueDiaryDungeonRsp               = 8362
	TryInterruptRogueDiaryDungeonReq               = 8617
	TryInterruptRogueDiaryDungeonRsp               = 8903
	UgcNotify                                      = 6341
	UnfreezeGroupLimitNotify                       = 3220
	UnionCmdNotify                                 = 5
	UnlockAvatarTalentReq                          = 1072
	UnlockAvatarTalentRsp                          = 1098
	UnlockCoopChapterReq                           = 1970
	UnlockCoopChapterRsp                           = 1995
	UnlockNameCardNotify                           = 4006
	UnlockPersonalLineReq                          = 449
	UnlockPersonalLineRsp                          = 491
	UnlockTransPointReq                            = 3035
	UnlockTransPointRsp                            = 3426
	UnlockedFurnitureFormulaDataNotify             = 4846
	UnlockedFurnitureSuiteDataNotify               = 4454
	UnmarkEntityInMinMapNotify                     = 219
	UpdateAbilityCreatedMovingPlatformNotify       = 881
	UpdatePS4BlockListReq                          = 4046
	UpdatePS4BlockListRsp                          = 4041
	UpdatePS4FriendListNotify                      = 4039
	UpdatePS4FriendListReq                         = 4089
	UpdatePS4FriendListRsp                         = 4059
	UpdatePlayerShowAvatarListReq                  = 4067
	UpdatePlayerShowAvatarListRsp                  = 4058
	UpdatePlayerShowNameCardListReq                = 4002
	UpdatePlayerShowNameCardListRsp                = 4019
	UpdateRedPointNotify                           = 93
	UpdateReunionWatcherNotify                     = 5091
	UpdateSalvageBundleMarkReq                     = 8967
	UpdateSalvageBundleMarkRsp                     = 8459
	UpgradeRoguelikeShikigamiReq                   = 8151
	UpgradeRoguelikeShikigamiRsp                   = 8966
	UseItemReq                                     = 690
	UseItemRsp                                     = 673
	UseMiracleRingReq                              = 5226
	UseMiracleRingRsp                              = 5218
	UseWidgetCreateGadgetReq                       = 4293
	UseWidgetCreateGadgetRsp                       = 4290
	UseWidgetRetractGadgetReq                      = 4286
	UseWidgetRetractGadgetRsp                      = 4261
	VehicleInteractReq                             = 865
	VehicleInteractRsp                             = 804
	VehicleStaminaNotify                           = 834
	ViewCodexReq                                   = 4202
	ViewCodexRsp                                   = 4201
	ViewLanternProjectionLevelTipsReq              = 8758
	ViewLanternProjectionLevelTipsRsp              = 8411
	ViewLanternProjectionTipsReq                   = 8218
	ViewLanternProjectionTipsRsp                   = 8590
	VintageCampGroupBundleRegisterNotify           = 24244
	VintageCampStageFinishNotify                   = 22830
	VintageDecorateBoothReq                        = 20846
	VintageDecorateBoothRsp                        = 20993
	VintageHuntingStartGalleryReq                  = 21780
	VintageHuntingStartGalleryRsp                  = 21951
	VintageMarketDeliverItemReq                    = 23141
	VintageMarketDeliverItemRsp                    = 22181
	VintageMarketDividendFinishNotify              = 23147
	VintageMarketFinishStorePlayReq                = 20676
	VintageMarketFinishStorePlayRsp                = 23462
	VintageMarketNpcEventFinishNotify              = 24201
	VintageMarketStartStorePlayReq                 = 22864
	VintageMarketStartStorePlayRsp                 = 22130
	VintageMarketStoreChooseStrategyReq            = 21248
	VintageMarketStoreChooseStrategyRsp            = 24860
	VintageMarketStoreUnlockSlotReq                = 20626
	VintageMarketStoreUnlockSlotRsp                = 20733
	VintageMarketStoreViewStrategyReq              = 21700
	VintageMarketStoreViewStrategyRsp              = 21814
	VintagePresentFinishNoify                      = 24142
	VintagePresentFinishNotify                     = 20086
	WatcherAllDataNotify                           = 2272
	WatcherChangeNotify                            = 2298
	WatcherEventNotify                             = 2212
	WatcherEventStageNotify                        = 2207
	WatcherEventTypeNotify                         = 2235
	WaterSpritePhaseFinishNotify                   = 2025
	WeaponAwakenReq                                = 695
	WeaponAwakenRsp                                = 606
	WeaponPromoteReq                               = 622
	WeaponPromoteRsp                               = 665
	WeaponUpgradeReq                               = 639
	WeaponUpgradeRsp                               = 653
	WearEquipReq                                   = 697
	WearEquipRsp                                   = 681
	WidgetActiveChangeNotify                       = 4280
	WidgetCaptureAnimalReq                         = 4256
	WidgetCaptureAnimalRsp                         = 4289
	WidgetCoolDownNotify                           = 4295
	WidgetDoBagReq                                 = 4255
	WidgetDoBagRsp                                 = 4296
	WidgetGadgetAllDataNotify                      = 4284
	WidgetGadgetDataNotify                         = 4266
	WidgetGadgetDestroyNotify                      = 4274
	WidgetQuickHitTreeReq                          = 3345
	WidgetQuickHitTreeRsp                          = 3336
	WidgetReportReq                                = 4291
	WidgetReportRsp                                = 4292
	WidgetSlotChangeNotify                         = 4267
	WidgetUpdateExtraCDReq                         = 5960
	WidgetUpdateExtraCDRsp                         = 6056
	WidgetUseAttachAbilityGroupChangeNotify        = 4258
	WindFieldGalleryChallengeInfoNotify            = 5563
	WindFieldGalleryInfoNotify                     = 5526
	WindFieldRestartDungeonReq                     = 20731
	WindFieldRestartDungeonRsp                     = 24712
	WindSeedClientNotify                           = 1199
	WinterCampAcceptAllGiveItemReq                 = 9000
	WinterCampAcceptAllGiveItemRsp                 = 8626
	WinterCampAcceptGiveItemReq                    = 8387
	WinterCampAcceptGiveItemRsp                    = 8185
	WinterCampEditSnowmanCombinationReq            = 8144
	WinterCampEditSnowmanCombinationRsp            = 8142
	WinterCampGetCanGiveFriendItemReq              = 8964
	WinterCampGetCanGiveFriendItemRsp              = 8357
	WinterCampGetFriendWishListReq                 = 8946
	WinterCampGetFriendWishListRsp                 = 8937
	WinterCampGetRecvItemListReq                   = 8143
	WinterCampGetRecvItemListRsp                   = 8423
	WinterCampGiveFriendItemReq                    = 8572
	WinterCampGiveFriendItemRsp                    = 8264
	WinterCampRaceScoreNotify                      = 8149
	WinterCampRecvItemNotify                       = 8580
	WinterCampSetWishListReq                       = 8753
	WinterCampSetWishListRsp                       = 8281
	WinterCampStageInfoChangeNotify                = 8154
	WinterCampTakeBattleRewardReq                  = 8401
	WinterCampTakeBattleRewardRsp                  = 8153
	WinterCampTakeExploreRewardReq                 = 8607
	WinterCampTakeExploreRewardRsp                 = 8978
	WinterCampTriathlonRestartReq                  = 8844
	WinterCampTriathlonRestartRsp                  = 8569
	WinterCampTriathlonSettleNotify                = 8342
	WorktopOptionNotify                            = 835
	WorldAllRoutineTypeNotify                      = 3518
	WorldChestOpenNotify                           = 3295
	WorldDataNotify                                = 3308
	WorldOwnerBlossomBriefInfoNotify               = 2735
	WorldOwnerBlossomScheduleInfoNotify            = 2707
	WorldOwnerDailyTaskNotify                      = 102
	WorldPlayerDieNotify                           = 285
	WorldPlayerInfoNotify                          = 3116
	WorldPlayerLocationNotify                      = 258
	WorldPlayerRTTNotify                           = 22
	WorldPlayerReviveReq                           = 225
	WorldPlayerReviveRsp                           = 278
	WorldRoutineChangeNotify                       = 3507
	WorldRoutineTypeCloseNotify                    = 3502
	WorldRoutineTypeRefreshNotify                  = 3525
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(AbilityChangeNotify, func() any { return new(proto.AbilityChangeNotify) })
	c.regMsg(AbilityInvocationFailNotify, func() any { return new(proto.AbilityInvocationFailNotify) })
	c.regMsg(AbilityInvocationFixedNotify, func() any { return new(proto.AbilityInvocationFixedNotify) })
	c.regMsg(AbilityInvocationsNotify, func() any { return new(proto.AbilityInvocationsNotify) })
	c.regMsg(AcceptCityReputationRequestReq, func() any { return new(proto.AcceptCityReputationRequestReq) })
	c.regMsg(AcceptCityReputationRequestRsp, func() any { return new(proto.AcceptCityReputationRequestRsp) })
	c.regMsg(AchievementAllDataNotify, func() any { return new(proto.AchievementAllDataNotify) })
	c.regMsg(AchievementUpdateNotify, func() any { return new(proto.AchievementUpdateNotify) })
	c.regMsg(ActivityAcceptAllGiveGiftReq, func() any { return new(proto.ActivityAcceptAllGiveGiftReq) })
	c.regMsg(ActivityAcceptAllGiveGiftRsp, func() any { return new(proto.ActivityAcceptAllGiveGiftRsp) })
	c.regMsg(ActivityAcceptGiveGiftReq, func() any { return new(proto.ActivityAcceptGiveGiftReq) })
	c.regMsg(ActivityAcceptGiveGiftRsp, func() any { return new(proto.ActivityAcceptGiveGiftRsp) })
	c.regMsg(ActivityBannerClearReq, func() any { return new(proto.ActivityBannerClearReq) })
	c.regMsg(ActivityBannerClearRsp, func() any { return new(proto.ActivityBannerClearRsp) })
	c.regMsg(ActivityBannerNotify, func() any { return new(proto.ActivityBannerNotify) })
	c.regMsg(ActivityCoinInfoNotify, func() any { return new(proto.ActivityCoinInfoNotify) })
	c.regMsg(ActivityCondStateChangeNotify, func() any { return new(proto.ActivityCondStateChangeNotify) })
	c.regMsg(ActivityDisableTransferPointInteractionNotify, func() any { return new(proto.ActivityDisableTransferPointInteractionNotify) })
	c.regMsg(ActivityGetCanGiveFriendGiftReq, func() any { return new(proto.ActivityGetCanGiveFriendGiftReq) })
	c.regMsg(ActivityGetCanGiveFriendGiftRsp, func() any { return new(proto.ActivityGetCanGiveFriendGiftRsp) })
	c.regMsg(ActivityGetFriendGiftWishListReq, func() any { return new(proto.ActivityGetFriendGiftWishListReq) })
	c.regMsg(ActivityGetFriendGiftWishListRsp, func() any { return new(proto.ActivityGetFriendGiftWishListRsp) })
	c.regMsg(ActivityGetRecvGiftListReq, func() any { return new(proto.ActivityGetRecvGiftListReq) })
	c.regMsg(ActivityGetRecvGiftListRsp, func() any { return new(proto.ActivityGetRecvGiftListRsp) })
	c.regMsg(ActivityGiveFriendGiftReq, func() any { return new(proto.ActivityGiveFriendGiftReq) })
	c.regMsg(ActivityGiveFriendGiftRsp, func() any { return new(proto.ActivityGiveFriendGiftRsp) })
	c.regMsg(ActivityHaveRecvGiftNotify, func() any { return new(proto.ActivityHaveRecvGiftNotify) })
	c.regMsg(ActivityInfoNotify, func() any { return new(proto.ActivityInfoNotify) })
	c.regMsg(ActivityPlayOpenAnimNotify, func() any { return new(proto.ActivityPlayOpenAnimNotify) })
	c.regMsg(ActivityPushTipsInfoNotify, func() any { return new(proto.ActivityPushTipsInfoNotify) })
	c.regMsg(ActivityReadPushTipsReq, func() any { return new(proto.ActivityReadPushTipsReq) })
	c.regMsg(ActivityReadPushTipsRsp, func() any { return new(proto.ActivityReadPushTipsRsp) })
	c.regMsg(ActivitySaleChangeNotify, func() any { return new(proto.ActivitySaleChangeNotify) })
	c.regMsg(ActivityScheduleInfoNotify, func() any { return new(proto.ActivityScheduleInfoNotify) })
	c.regMsg(ActivitySelectAvatarCardReq, func() any { return new(proto.ActivitySelectAvatarCardReq) })
	c.regMsg(ActivitySelectAvatarCardRsp, func() any { return new(proto.ActivitySelectAvatarCardRsp) })
	c.regMsg(ActivitySetGiftWishReq, func() any { return new(proto.ActivitySetGiftWishReq) })
	c.regMsg(ActivitySetGiftWishRsp, func() any { return new(proto.ActivitySetGiftWishRsp) })
	c.regMsg(ActivityTakeAllScoreRewardReq, func() any { return new(proto.ActivityTakeAllScoreRewardReq) })
	c.regMsg(ActivityTakeAllScoreRewardRsp, func() any { return new(proto.ActivityTakeAllScoreRewardRsp) })
	c.regMsg(ActivityTakeScoreRewardReq, func() any { return new(proto.ActivityTakeScoreRewardReq) })
	c.regMsg(ActivityTakeScoreRewardRsp, func() any { return new(proto.ActivityTakeScoreRewardRsp) })
	c.regMsg(ActivityTakeWatcherRewardBatchReq, func() any { return new(proto.ActivityTakeWatcherRewardBatchReq) })
	c.regMsg(ActivityTakeWatcherRewardBatchRsp, func() any { return new(proto.ActivityTakeWatcherRewardBatchRsp) })
	c.regMsg(ActivityTakeWatcherRewardReq, func() any { return new(proto.ActivityTakeWatcherRewardReq) })
	c.regMsg(ActivityTakeWatcherRewardRsp, func() any { return new(proto.ActivityTakeWatcherRewardRsp) })
	c.regMsg(ActivityUpdateWatcherNotify, func() any { return new(proto.ActivityUpdateWatcherNotify) })
	c.regMsg(AddAranaraCollectionNotify, func() any { return new(proto.AddAranaraCollectionNotify) })
	c.regMsg(AddBackupAvatarTeamReq, func() any { return new(proto.AddBackupAvatarTeamReq) })
	c.regMsg(AddBackupAvatarTeamRsp, func() any { return new(proto.AddBackupAvatarTeamRsp) })
	c.regMsg(AddBlacklistReq, func() any { return new(proto.AddBlacklistReq) })
	c.regMsg(AddBlacklistRsp, func() any { return new(proto.AddBlacklistRsp) })
	c.regMsg(AddFriendNotify, func() any { return new(proto.AddFriendNotify) })
	c.regMsg(AddNoGachaAvatarCardNotify, func() any { return new(proto.AddNoGachaAvatarCardNotify) })
	c.regMsg(AddQuestContentProgressReq, func() any { return new(proto.AddQuestContentProgressReq) })
	c.regMsg(AddQuestContentProgressRsp, func() any { return new(proto.AddQuestContentProgressRsp) })
	c.regMsg(AddRandTaskInfoNotify, func() any { return new(proto.AddRandTaskInfoNotify) })
	c.regMsg(AddSeenMonsterNotify, func() any { return new(proto.AddSeenMonsterNotify) })
	c.regMsg(AdjustWorldLevelReq, func() any { return new(proto.AdjustWorldLevelReq) })
	c.regMsg(AdjustWorldLevelRsp, func() any { return new(proto.AdjustWorldLevelRsp) })
	c.regMsg(AllCoopInfoNotify, func() any { return new(proto.AllCoopInfoNotify) })
	c.regMsg(AllMarkPointNotify, func() any { return new(proto.AllMarkPointNotify) })
	c.regMsg(AllSeenMonsterNotify, func() any { return new(proto.AllSeenMonsterNotify) })
	c.regMsg(AllShareCDDataNotify, func() any { return new(proto.AllShareCDDataNotify) })
	c.regMsg(AllWidgetBackgroundActiveStateNotify, func() any { return new(proto.AllWidgetBackgroundActiveStateNotify) })
	c.regMsg(AllWidgetDataNotify, func() any { return new(proto.AllWidgetDataNotify) })
	c.regMsg(AnchorPointDataNotify, func() any { return new(proto.AnchorPointDataNotify) })
	c.regMsg(AnchorPointOpReq, func() any { return new(proto.AnchorPointOpReq) })
	c.regMsg(AnchorPointOpRsp, func() any { return new(proto.AnchorPointOpRsp) })
	c.regMsg(AnimatorForceSetAirMoveNotify, func() any { return new(proto.AnimatorForceSetAirMoveNotify) })
	c.regMsg(AntiAddictNotify, func() any { return new(proto.AntiAddictNotify) })
	c.regMsg(AranaraCollectionDataNotify, func() any { return new(proto.AranaraCollectionDataNotify) })
	c.regMsg(AreaPlayInfoNotify, func() any { return new(proto.AreaPlayInfoNotify) })
	c.regMsg(ArenaChallengeFinishNotify, func() any { return new(proto.ArenaChallengeFinishNotify) })
	c.regMsg(AskAddFriendNotify, func() any { return new(proto.AskAddFriendNotify) })
	c.regMsg(AskAddFriendReq, func() any { return new(proto.AskAddFriendReq) })
	c.regMsg(AskAddFriendRsp, func() any { return new(proto.AskAddFriendRsp) })
	c.regMsg(AssociateInferenceWordReq, func() any { return new(proto.AssociateInferenceWordReq) })
	c.regMsg(AssociateInferenceWordRsp, func() any { return new(proto.AssociateInferenceWordRsp) })
	c.regMsg(AsterLargeInfoNotify, func() any { return new(proto.AsterLargeInfoNotify) })
	c.regMsg(AsterLittleInfoNotify, func() any { return new(proto.AsterLittleInfoNotify) })
	c.regMsg(AsterMidCampInfoNotify, func() any { return new(proto.AsterMidCampInfoNotify) })
	c.regMsg(AsterMidInfoNotify, func() any { return new(proto.AsterMidInfoNotify) })
	c.regMsg(AsterMiscInfoNotify, func() any { return new(proto.AsterMiscInfoNotify) })
	c.regMsg(AsterProgressInfoNotify, func() any { return new(proto.AsterProgressInfoNotify) })
	c.regMsg(AvatarAddNotify, func() any { return new(proto.AvatarAddNotify) })
	c.regMsg(AvatarBuffAddNotify, func() any { return new(proto.AvatarBuffAddNotify) })
	c.regMsg(AvatarBuffDelNotify, func() any { return new(proto.AvatarBuffDelNotify) })
	c.regMsg(AvatarCardChangeReq, func() any { return new(proto.AvatarCardChangeReq) })
	c.regMsg(AvatarCardChangeRsp, func() any { return new(proto.AvatarCardChangeRsp) })
	c.regMsg(AvatarChangeAnimHashReq, func() any { return new(proto.AvatarChangeAnimHashReq) })
	c.regMsg(AvatarChangeAnimHashRsp, func() any { return new(proto.AvatarChangeAnimHashRsp) })
	c.regMsg(AvatarChangeCostumeNotify, func() any { return new(proto.AvatarChangeCostumeNotify) })
	c.regMsg(AvatarChangeCostumeReq, func() any { return new(proto.AvatarChangeCostumeReq) })
	c.regMsg(AvatarChangeCostumeRsp, func() any { return new(proto.AvatarChangeCostumeRsp) })
	c.regMsg(AvatarChangeElementTypeReq, func() any { return new(proto.AvatarChangeElementTypeReq) })
	c.regMsg(AvatarChangeElementTypeRsp, func() any { return new(proto.AvatarChangeElementTypeRsp) })
	c.regMsg(AvatarDataNotify, func() any { return new(proto.AvatarDataNotify) })
	c.regMsg(AvatarDelNotify, func() any { return new(proto.AvatarDelNotify) })
	c.regMsg(AvatarDieAnimationEndReq, func() any { return new(proto.AvatarDieAnimationEndReq) })
	c.regMsg(AvatarDieAnimationEndRsp, func() any { return new(proto.AvatarDieAnimationEndRsp) })
	c.regMsg(AvatarEnterElementViewNotify, func() any { return new(proto.AvatarEnterElementViewNotify) })
	c.regMsg(AvatarEquipAffixStartNotify, func() any { return new(proto.AvatarEquipAffixStartNotify) })
	c.regMsg(AvatarEquipChangeNotify, func() any { return new(proto.AvatarEquipChangeNotify) })
	c.regMsg(AvatarExpeditionAllDataReq, func() any { return new(proto.AvatarExpeditionAllDataReq) })
	c.regMsg(AvatarExpeditionAllDataRsp, func() any { return new(proto.AvatarExpeditionAllDataRsp) })
	c.regMsg(AvatarExpeditionCallBackReq, func() any { return new(proto.AvatarExpeditionCallBackReq) })
	c.regMsg(AvatarExpeditionCallBackRsp, func() any { return new(proto.AvatarExpeditionCallBackRsp) })
	c.regMsg(AvatarExpeditionDataNotify, func() any { return new(proto.AvatarExpeditionDataNotify) })
	c.regMsg(AvatarExpeditionGetRewardReq, func() any { return new(proto.AvatarExpeditionGetRewardReq) })
	c.regMsg(AvatarExpeditionGetRewardRsp, func() any { return new(proto.AvatarExpeditionGetRewardRsp) })
	c.regMsg(AvatarExpeditionStartReq, func() any { return new(proto.AvatarExpeditionStartReq) })
	c.regMsg(AvatarExpeditionStartRsp, func() any { return new(proto.AvatarExpeditionStartRsp) })
	c.regMsg(AvatarFetterDataNotify, func() any { return new(proto.AvatarFetterDataNotify) })
	c.regMsg(AvatarFetterLevelRewardReq, func() any { return new(proto.AvatarFetterLevelRewardReq) })
	c.regMsg(AvatarFetterLevelRewardRsp, func() any { return new(proto.AvatarFetterLevelRewardRsp) })
	c.regMsg(AvatarFightPropNotify, func() any { return new(proto.AvatarFightPropNotify) })
	c.regMsg(AvatarFightPropUpdateNotify, func() any { return new(proto.AvatarFightPropUpdateNotify) })
	c.regMsg(AvatarFlycloakChangeNotify, func() any { return new(proto.AvatarFlycloakChangeNotify) })
	c.regMsg(AvatarFollowRouteNotify, func() any { return new(proto.AvatarFollowRouteNotify) })
	c.regMsg(AvatarGainCostumeNotify, func() any { return new(proto.AvatarGainCostumeNotify) })
	c.regMsg(AvatarGainFlycloakNotify, func() any { return new(proto.AvatarGainFlycloakNotify) })
	c.regMsg(AvatarLifeStateChangeNotify, func() any { return new(proto.AvatarLifeStateChangeNotify) })
	c.regMsg(AvatarPromoteGetRewardReq, func() any { return new(proto.AvatarPromoteGetRewardReq) })
	c.regMsg(AvatarPromoteGetRewardRsp, func() any { return new(proto.AvatarPromoteGetRewardRsp) })
	c.regMsg(AvatarPromoteReq, func() any { return new(proto.AvatarPromoteReq) })
	c.regMsg(AvatarPromoteRsp, func() any { return new(proto.AvatarPromoteRsp) })
	c.regMsg(AvatarPropChangeReasonNotify, func() any { return new(proto.AvatarPropChangeReasonNotify) })
	c.regMsg(AvatarPropNotify, func() any { return new(proto.AvatarPropNotify) })
	c.regMsg(AvatarSatiationDataNotify, func() any { return new(proto.AvatarSatiationDataNotify) })
	c.regMsg(AvatarSkillChangeNotify, func() any { return new(proto.AvatarSkillChangeNotify) })
	c.regMsg(AvatarSkillDepotChangeNotify, func() any { return new(proto.AvatarSkillDepotChangeNotify) })
	c.regMsg(AvatarSkillInfoNotify, func() any { return new(proto.AvatarSkillInfoNotify) })
	c.regMsg(AvatarSkillMaxChargeCountNotify, func() any { return new(proto.AvatarSkillMaxChargeCountNotify) })
	c.regMsg(AvatarSkillUpgradeReq, func() any { return new(proto.AvatarSkillUpgradeReq) })
	c.regMsg(AvatarSkillUpgradeRsp, func() any { return new(proto.AvatarSkillUpgradeRsp) })
	c.regMsg(AvatarTeamAllDataNotify, func() any { return new(proto.AvatarTeamAllDataNotify) })
	c.regMsg(AvatarTeamUpdateNotify, func() any { return new(proto.AvatarTeamUpdateNotify) })
	c.regMsg(AvatarUnlockTalentNotify, func() any { return new(proto.AvatarUnlockTalentNotify) })
	c.regMsg(AvatarUpgradeReq, func() any { return new(proto.AvatarUpgradeReq) })
	c.regMsg(AvatarUpgradeRsp, func() any { return new(proto.AvatarUpgradeRsp) })
	c.regMsg(AvatarWearFlycloakReq, func() any { return new(proto.AvatarWearFlycloakReq) })
	c.regMsg(AvatarWearFlycloakRsp, func() any { return new(proto.AvatarWearFlycloakRsp) })
	c.regMsg(BackMyWorldReq, func() any { return new(proto.BackMyWorldReq) })
	c.regMsg(BackMyWorldRsp, func() any { return new(proto.BackMyWorldRsp) })
	c.regMsg(BackPlayCustomDungeonOfficialReq, func() any { return new(proto.BackPlayCustomDungeonOfficialReq) })
	c.regMsg(BackPlayCustomDungeonOfficialRsp, func() any { return new(proto.BackPlayCustomDungeonOfficialRsp) })
	c.regMsg(BackRebornGalleryReq, func() any { return new(proto.BackRebornGalleryReq) })
	c.regMsg(BackRebornGalleryRsp, func() any { return new(proto.BackRebornGalleryRsp) })
	c.regMsg(BargainOfferPriceReq, func() any { return new(proto.BargainOfferPriceReq) })
	c.regMsg(BargainOfferPriceRsp, func() any { return new(proto.BargainOfferPriceRsp) })
	c.regMsg(BargainStartNotify, func() any { return new(proto.BargainStartNotify) })
	c.regMsg(BargainTerminateNotify, func() any { return new(proto.BargainTerminateNotify) })
	c.regMsg(BartenderCancelLevelReq, func() any { return new(proto.BartenderCancelLevelReq) })
	c.regMsg(BartenderCancelLevelRsp, func() any { return new(proto.BartenderCancelLevelRsp) })
	c.regMsg(BartenderCancelOrderReq, func() any { return new(proto.BartenderCancelOrderReq) })
	c.regMsg(BartenderCancelOrderRsp, func() any { return new(proto.BartenderCancelOrderRsp) })
	c.regMsg(BartenderCompleteOrderReq, func() any { return new(proto.BartenderCompleteOrderReq) })
	c.regMsg(BartenderCompleteOrderRsp, func() any { return new(proto.BartenderCompleteOrderRsp) })
	c.regMsg(BartenderFinishLevelReq, func() any { return new(proto.BartenderFinishLevelReq) })
	c.regMsg(BartenderFinishLevelRsp, func() any { return new(proto.BartenderFinishLevelRsp) })
	c.regMsg(BartenderGetFormulaReq, func() any { return new(proto.BartenderGetFormulaReq) })
	c.regMsg(BartenderGetFormulaRsp, func() any { return new(proto.BartenderGetFormulaRsp) })
	c.regMsg(BartenderLevelProgressNotify, func() any { return new(proto.BartenderLevelProgressNotify) })
	c.regMsg(BartenderStartLevelReq, func() any { return new(proto.BartenderStartLevelReq) })
	c.regMsg(BartenderStartLevelRsp, func() any { return new(proto.BartenderStartLevelRsp) })
	c.regMsg(BattlePassAllDataNotify, func() any { return new(proto.BattlePassAllDataNotify) })
	c.regMsg(BattlePassBuySuccNotify, func() any { return new(proto.BattlePassBuySuccNotify) })
	c.regMsg(BattlePassCurScheduleUpdateNotify, func() any { return new(proto.BattlePassCurScheduleUpdateNotify) })
	c.regMsg(BattlePassMissionDelNotify, func() any { return new(proto.BattlePassMissionDelNotify) })
	c.regMsg(BattlePassMissionUpdateNotify, func() any { return new(proto.BattlePassMissionUpdateNotify) })
	c.regMsg(BeginCameraSceneLookNotify, func() any { return new(proto.BeginCameraSceneLookNotify) })
	c.regMsg(BeginCameraSceneLookWithTemplateNotify, func() any { return new(proto.BeginCameraSceneLookWithTemplateNotify) })
	c.regMsg(BigTalentPointConvertReq, func() any { return new(proto.BigTalentPointConvertReq) })
	c.regMsg(BigTalentPointConvertRsp, func() any { return new(proto.BigTalentPointConvertRsp) })
	c.regMsg(BlessingAcceptAllGivePicReq, func() any { return new(proto.BlessingAcceptAllGivePicReq) })
	c.regMsg(BlessingAcceptAllGivePicRsp, func() any { return new(proto.BlessingAcceptAllGivePicRsp) })
	c.regMsg(BlessingAcceptGivePicReq, func() any { return new(proto.BlessingAcceptGivePicReq) })
	c.regMsg(BlessingAcceptGivePicRsp, func() any { return new(proto.BlessingAcceptGivePicRsp) })
	c.regMsg(BlessingGetAllRecvPicRecordListReq, func() any { return new(proto.BlessingGetAllRecvPicRecordListReq) })
	c.regMsg(BlessingGetAllRecvPicRecordListRsp, func() any { return new(proto.BlessingGetAllRecvPicRecordListRsp) })
	c.regMsg(BlessingGetFriendPicListReq, func() any { return new(proto.BlessingGetFriendPicListReq) })
	c.regMsg(BlessingGetFriendPicListRsp, func() any { return new(proto.BlessingGetFriendPicListRsp) })
	c.regMsg(BlessingGiveFriendPicReq, func() any { return new(proto.BlessingGiveFriendPicReq) })
	c.regMsg(BlessingGiveFriendPicRsp, func() any { return new(proto.BlessingGiveFriendPicRsp) })
	c.regMsg(BlessingRecvFriendPicNotify, func() any { return new(proto.BlessingRecvFriendPicNotify) })
	c.regMsg(BlessingRedeemRewardReq, func() any { return new(proto.BlessingRedeemRewardReq) })
	c.regMsg(BlessingRedeemRewardRsp, func() any { return new(proto.BlessingRedeemRewardRsp) })
	c.regMsg(BlessingScanReq, func() any { return new(proto.BlessingScanReq) })
	c.regMsg(BlessingScanRsp, func() any { return new(proto.BlessingScanRsp) })
	c.regMsg(BlitzRushParkourRestartReq, func() any { return new(proto.BlitzRushParkourRestartReq) })
	c.regMsg(BlitzRushParkourRestartRsp, func() any { return new(proto.BlitzRushParkourRestartRsp) })
	c.regMsg(BlossomBriefInfoNotify, func() any { return new(proto.BlossomBriefInfoNotify) })
	c.regMsg(BlossomChestCreateNotify, func() any { return new(proto.BlossomChestCreateNotify) })
	c.regMsg(BlossomChestInfoNotify, func() any { return new(proto.BlossomChestInfoNotify) })
	c.regMsg(BonusActivityInfoReq, func() any { return new(proto.BonusActivityInfoReq) })
	c.regMsg(BonusActivityInfoRsp, func() any { return new(proto.BonusActivityInfoRsp) })
	c.regMsg(BonusActivityUpdateNotify, func() any { return new(proto.BonusActivityUpdateNotify) })
	c.regMsg(BossChestActivateNotify, func() any { return new(proto.BossChestActivateNotify) })
	c.regMsg(BounceConjuringSettleNotify, func() any { return new(proto.BounceConjuringSettleNotify) })
	c.regMsg(BuoyantCombatSettleNotify, func() any { return new(proto.BuoyantCombatSettleNotify) })
	c.regMsg(BuyBattlePassLevelReq, func() any { return new(proto.BuyBattlePassLevelReq) })
	c.regMsg(BuyBattlePassLevelRsp, func() any { return new(proto.BuyBattlePassLevelRsp) })
	c.regMsg(BuyGoodsReq, func() any { return new(proto.BuyGoodsReq) })
	c.regMsg(BuyGoodsRsp, func() any { return new(proto.BuyGoodsRsp) })
	c.regMsg(BuyResinReq, func() any { return new(proto.BuyResinReq) })
	c.regMsg(BuyResinRsp, func() any { return new(proto.BuyResinRsp) })
	c.regMsg(CalcWeaponUpgradeReturnItemsReq, func() any { return new(proto.CalcWeaponUpgradeReturnItemsReq) })
	c.regMsg(CalcWeaponUpgradeReturnItemsRsp, func() any { return new(proto.CalcWeaponUpgradeReturnItemsRsp) })
	c.regMsg(CanUseSkillNotify, func() any { return new(proto.CanUseSkillNotify) })
	c.regMsg(CancelCityReputationRequestReq, func() any { return new(proto.CancelCityReputationRequestReq) })
	c.regMsg(CancelCityReputationRequestRsp, func() any { return new(proto.CancelCityReputationRequestRsp) })
	c.regMsg(CancelCoopTaskReq, func() any { return new(proto.CancelCoopTaskReq) })
	c.regMsg(CancelCoopTaskRsp, func() any { return new(proto.CancelCoopTaskRsp) })
	c.regMsg(CancelFinishParentQuestNotify, func() any { return new(proto.CancelFinishParentQuestNotify) })
	c.regMsg(CardProductRewardNotify, func() any { return new(proto.CardProductRewardNotify) })
	c.regMsg(CataLogFinishedGlobalWatcherAllDataNotify, func() any { return new(proto.CataLogFinishedGlobalWatcherAllDataNotify) })
	c.regMsg(CataLogNewFinishedGlobalWatcherNotify, func() any { return new(proto.CataLogNewFinishedGlobalWatcherNotify) })
	c.regMsg(ChallengeDataNotify, func() any { return new(proto.ChallengeDataNotify) })
	c.regMsg(ChallengeRecordNotify, func() any { return new(proto.ChallengeRecordNotify) })
	c.regMsg(ChangeAvatarReq, func() any { return new(proto.ChangeAvatarReq) })
	c.regMsg(ChangeAvatarRsp, func() any { return new(proto.ChangeAvatarRsp) })
	c.regMsg(ChangeCustomDungeonRoomReq, func() any { return new(proto.ChangeCustomDungeonRoomReq) })
	c.regMsg(ChangeCustomDungeonRoomRsp, func() any { return new(proto.ChangeCustomDungeonRoomRsp) })
	c.regMsg(ChangeGameTimeReq, func() any { return new(proto.ChangeGameTimeReq) })
	c.regMsg(ChangeGameTimeRsp, func() any { return new(proto.ChangeGameTimeRsp) })
	c.regMsg(ChangeMailStarNotify, func() any { return new(proto.ChangeMailStarNotify) })
	c.regMsg(ChangeMpTeamAvatarReq, func() any { return new(proto.ChangeMpTeamAvatarReq) })
	c.regMsg(ChangeMpTeamAvatarRsp, func() any { return new(proto.ChangeMpTeamAvatarRsp) })
	c.regMsg(ChangeServerGlobalValueNotify, func() any { return new(proto.ChangeServerGlobalValueNotify) })
	c.regMsg(ChangeTeamNameReq, func() any { return new(proto.ChangeTeamNameReq) })
	c.regMsg(ChangeTeamNameRsp, func() any { return new(proto.ChangeTeamNameRsp) })
	c.regMsg(ChangeWidgetBackgroundActiveStateReq, func() any { return new(proto.ChangeWidgetBackgroundActiveStateReq) })
	c.regMsg(ChangeWidgetBackgroundActiveStateRsp, func() any { return new(proto.ChangeWidgetBackgroundActiveStateRsp) })
	c.regMsg(ChangeWorldToSingleModeNotify, func() any { return new(proto.ChangeWorldToSingleModeNotify) })
	c.regMsg(ChangeWorldToSingleModeReq, func() any { return new(proto.ChangeWorldToSingleModeReq) })
	c.regMsg(ChangeWorldToSingleModeRsp, func() any { return new(proto.ChangeWorldToSingleModeRsp) })
	c.regMsg(ChannelerSlabStageActiveChallengeIndexNotify, func() any { return new(proto.ChannelerSlabStageActiveChallengeIndexNotify) })
	c.regMsg(ChapterStateNotify, func() any { return new(proto.ChapterStateNotify) })
	c.regMsg(CharAmusementSettleNotify, func() any { return new(proto.CharAmusementSettleNotify) })
	c.regMsg(ChatChannelDataNotify, func() any { return new(proto.ChatChannelDataNotify) })
	c.regMsg(ChatChannelUpdateNotify, func() any { return new(proto.ChatChannelUpdateNotify) })
	c.regMsg(ChatHistoryNotify, func() any { return new(proto.ChatHistoryNotify) })
	c.regMsg(CheckAddItemExceedLimitNotify, func() any { return new(proto.CheckAddItemExceedLimitNotify) })
	c.regMsg(CheckGroupReplacedReq, func() any { return new(proto.CheckGroupReplacedReq) })
	c.regMsg(CheckGroupReplacedRsp, func() any { return new(proto.CheckGroupReplacedRsp) })
	c.regMsg(CheckSegmentCRCNotify, func() any { return new(proto.CheckSegmentCRCNotify) })
	c.regMsg(CheckSegmentCRCReq, func() any { return new(proto.CheckSegmentCRCReq) })
	c.regMsg(CheckUgcStateReq, func() any { return new(proto.CheckUgcStateReq) })
	c.regMsg(CheckUgcStateRsp, func() any { return new(proto.CheckUgcStateRsp) })
	c.regMsg(CheckUgcUpdateReq, func() any { return new(proto.CheckUgcUpdateReq) })
	c.regMsg(CheckUgcUpdateRsp, func() any { return new(proto.CheckUgcUpdateRsp) })
	c.regMsg(ChessEscapedMonstersNotify, func() any { return new(proto.ChessEscapedMonstersNotify) })
	c.regMsg(ChessLeftMonstersNotify, func() any { return new(proto.ChessLeftMonstersNotify) })
	c.regMsg(ChessManualRefreshCardsReq, func() any { return new(proto.ChessManualRefreshCardsReq) })
	c.regMsg(ChessManualRefreshCardsRsp, func() any { return new(proto.ChessManualRefreshCardsRsp) })
	c.regMsg(ChessPickCardNotify, func() any { return new(proto.ChessPickCardNotify) })
	c.regMsg(ChessPickCardReq, func() any { return new(proto.ChessPickCardReq) })
	c.regMsg(ChessPickCardRsp, func() any { return new(proto.ChessPickCardRsp) })
	c.regMsg(ChessPlayerInfoNotify, func() any { return new(proto.ChessPlayerInfoNotify) })
	c.regMsg(ChessSelectedCardsNotify, func() any { return new(proto.ChessSelectedCardsNotify) })
	c.regMsg(ChooseCurAvatarTeamReq, func() any { return new(proto.ChooseCurAvatarTeamReq) })
	c.regMsg(ChooseCurAvatarTeamRsp, func() any { return new(proto.ChooseCurAvatarTeamRsp) })
	c.regMsg(CityReputationDataNotify, func() any { return new(proto.CityReputationDataNotify) })
	c.regMsg(CityReputationLevelupNotify, func() any { return new(proto.CityReputationLevelupNotify) })
	c.regMsg(ClearRoguelikeCurseNotify, func() any { return new(proto.ClearRoguelikeCurseNotify) })
	c.regMsg(ClientAIStateNotify, func() any { return new(proto.ClientAIStateNotify) })
	c.regMsg(ClientAbilitiesInitFinishCombineNotify, func() any { return new(proto.ClientAbilitiesInitFinishCombineNotify) })
	c.regMsg(ClientAbilityChangeNotify, func() any { return new(proto.ClientAbilityChangeNotify) })
	c.regMsg(ClientAbilityInitBeginNotify, func() any { return new(proto.ClientAbilityInitBeginNotify) })
	c.regMsg(ClientAbilityInitFinishNotify, func() any { return new(proto.ClientAbilityInitFinishNotify) })
	c.regMsg(ClientBulletCreateNotify, func() any { return new(proto.ClientBulletCreateNotify) })
	c.regMsg(ClientCollectorDataNotify, func() any { return new(proto.ClientCollectorDataNotify) })
	c.regMsg(ClientHashDebugNotify, func() any { return new(proto.ClientHashDebugNotify) })
	c.regMsg(ClientLoadingCostumeVerificationNotify, func() any { return new(proto.ClientLoadingCostumeVerificationNotify) })
	c.regMsg(ClientLockGameTimeNotify, func() any { return new(proto.ClientLockGameTimeNotify) })
	c.regMsg(ClientNewMailNotify, func() any { return new(proto.ClientNewMailNotify) })
	c.regMsg(ClientPauseNotify, func() any { return new(proto.ClientPauseNotify) })
	c.regMsg(ClientReconnectNotify, func() any { return new(proto.ClientReconnectNotify) })
	c.regMsg(ClientRemoveCombatEndModifierNotify, func() any { return new(proto.ClientRemoveCombatEndModifierNotify) })
	c.regMsg(ClientReportNotify, func() any { return new(proto.ClientReportNotify) })
	c.regMsg(ClientScriptEventNotify, func() any { return new(proto.ClientScriptEventNotify) })
	c.regMsg(ClientTransmitReq, func() any { return new(proto.ClientTransmitReq) })
	c.regMsg(ClientTransmitRsp, func() any { return new(proto.ClientTransmitRsp) })
	c.regMsg(ClientTriggerEventNotify, func() any { return new(proto.ClientTriggerEventNotify) })
	c.regMsg(CloseCommonTipsNotify, func() any { return new(proto.CloseCommonTipsNotify) })
	c.regMsg(ClosedItemNotify, func() any { return new(proto.ClosedItemNotify) })
	c.regMsg(CodexDataFullNotify, func() any { return new(proto.CodexDataFullNotify) })
	c.regMsg(CodexDataUpdateNotify, func() any { return new(proto.CodexDataUpdateNotify) })
	c.regMsg(CombatInvocationsNotify, func() any { return new(proto.CombatInvocationsNotify) })
	c.regMsg(CombineDataNotify, func() any { return new(proto.CombineDataNotify) })
	c.regMsg(CombineFormulaDataNotify, func() any { return new(proto.CombineFormulaDataNotify) })
	c.regMsg(CombineReq, func() any { return new(proto.CombineReq) })
	c.regMsg(CombineRsp, func() any { return new(proto.CombineRsp) })
	c.regMsg(CommonPlayerTipsNotify, func() any { return new(proto.CommonPlayerTipsNotify) })
	c.regMsg(CompoundDataNotify, func() any { return new(proto.CompoundDataNotify) })
	c.regMsg(CompoundUnlockNotify, func() any { return new(proto.CompoundUnlockNotify) })
	c.regMsg(CookDataNotify, func() any { return new(proto.CookDataNotify) })
	c.regMsg(CookGradeDataNotify, func() any { return new(proto.CookGradeDataNotify) })
	c.regMsg(CookRecipeDataNotify, func() any { return new(proto.CookRecipeDataNotify) })
	c.regMsg(CoopCgShowNotify, func() any { return new(proto.CoopCgShowNotify) })
	c.regMsg(CoopCgUpdateNotify, func() any { return new(proto.CoopCgUpdateNotify) })
	c.regMsg(CoopChapterUpdateNotify, func() any { return new(proto.CoopChapterUpdateNotify) })
	c.regMsg(CoopDataNotify, func() any { return new(proto.CoopDataNotify) })
	c.regMsg(CoopPointUpdateNotify, func() any { return new(proto.CoopPointUpdateNotify) })
	c.regMsg(CoopProgressUpdateNotify, func() any { return new(proto.CoopProgressUpdateNotify) })
	c.regMsg(CoopRewardUpdateNotify, func() any { return new(proto.CoopRewardUpdateNotify) })
	c.regMsg(CreateMassiveEntityNotify, func() any { return new(proto.CreateMassiveEntityNotify) })
	c.regMsg(CreateMassiveEntityReq, func() any { return new(proto.CreateMassiveEntityReq) })
	c.regMsg(CreateMassiveEntityRsp, func() any { return new(proto.CreateMassiveEntityRsp) })
	c.regMsg(CreateVehicleReq, func() any { return new(proto.CreateVehicleReq) })
	c.regMsg(CreateVehicleRsp, func() any { return new(proto.CreateVehicleRsp) })
	c.regMsg(CrystalLinkDungeonInfoNotify, func() any { return new(proto.CrystalLinkDungeonInfoNotify) })
	c.regMsg(CrystalLinkEnterDungeonReq, func() any { return new(proto.CrystalLinkEnterDungeonReq) })
	c.regMsg(CrystalLinkEnterDungeonRsp, func() any { return new(proto.CrystalLinkEnterDungeonRsp) })
	c.regMsg(CrystalLinkRestartDungeonReq, func() any { return new(proto.CrystalLinkRestartDungeonReq) })
	c.regMsg(CrystalLinkRestartDungeonRsp, func() any { return new(proto.CrystalLinkRestartDungeonRsp) })
	c.regMsg(CustomDungeonBattleRecordNotify, func() any { return new(proto.CustomDungeonBattleRecordNotify) })
	c.regMsg(CustomDungeonOfficialNotify, func() any { return new(proto.CustomDungeonOfficialNotify) })
	c.regMsg(CustomDungeonRecoverNotify, func() any { return new(proto.CustomDungeonRecoverNotify) })
	c.regMsg(CustomDungeonUpdateNotify, func() any { return new(proto.CustomDungeonUpdateNotify) })
	c.regMsg(CutSceneBeginNotify, func() any { return new(proto.CutSceneBeginNotify) })
	c.regMsg(CutSceneEndNotify, func() any { return new(proto.CutSceneEndNotify) })
	c.regMsg(CutSceneFinishNotify, func() any { return new(proto.CutSceneFinishNotify) })
	c.regMsg(DailyTaskDataNotify, func() any { return new(proto.DailyTaskDataNotify) })
	c.regMsg(DailyTaskFilterCityReq, func() any { return new(proto.DailyTaskFilterCityReq) })
	c.regMsg(DailyTaskFilterCityRsp, func() any { return new(proto.DailyTaskFilterCityRsp) })
	c.regMsg(DailyTaskProgressNotify, func() any { return new(proto.DailyTaskProgressNotify) })
	c.regMsg(DailyTaskScoreRewardNotify, func() any { return new(proto.DailyTaskScoreRewardNotify) })
	c.regMsg(DailyTaskUnlockedCitiesNotify, func() any { return new(proto.DailyTaskUnlockedCitiesNotify) })
	c.regMsg(DataResVersionNotify, func() any { return new(proto.DataResVersionNotify) })
	c.regMsg(DealAddFriendReq, func() any { return new(proto.DealAddFriendReq) })
	c.regMsg(DealAddFriendRsp, func() any { return new(proto.DealAddFriendRsp) })
	c.regMsg(DeathZoneInfoNotify, func() any { return new(proto.DeathZoneInfoNotify) })
	c.regMsg(DeathZoneObserveNotify, func() any { return new(proto.DeathZoneObserveNotify) })
	c.regMsg(DebugNotify, func() any { return new(proto.DebugNotify) })
	c.regMsg(DelBackupAvatarTeamReq, func() any { return new(proto.DelBackupAvatarTeamReq) })
	c.regMsg(DelBackupAvatarTeamRsp, func() any { return new(proto.DelBackupAvatarTeamRsp) })
	c.regMsg(DelMailReq, func() any { return new(proto.DelMailReq) })
	c.regMsg(DelMailRsp, func() any { return new(proto.DelMailRsp) })
	c.regMsg(DelScenePlayTeamEntityNotify, func() any { return new(proto.DelScenePlayTeamEntityNotify) })
	c.regMsg(DelTeamEntityNotify, func() any { return new(proto.DelTeamEntityNotify) })
	c.regMsg(DeleteFriendNotify, func() any { return new(proto.DeleteFriendNotify) })
	c.regMsg(DeleteFriendReq, func() any { return new(proto.DeleteFriendReq) })
	c.regMsg(DeleteFriendRsp, func() any { return new(proto.DeleteFriendRsp) })
	c.regMsg(DeshretObeliskChestInfoNotify, func() any { return new(proto.DeshretObeliskChestInfoNotify) })
	c.regMsg(DestroyMassiveEntityNotify, func() any { return new(proto.DestroyMassiveEntityNotify) })
	c.regMsg(DestroyMaterialReq, func() any { return new(proto.DestroyMaterialReq) })
	c.regMsg(DestroyMaterialRsp, func() any { return new(proto.DestroyMaterialRsp) })
	c.regMsg(DigActivityChangeGadgetStateReq, func() any { return new(proto.DigActivityChangeGadgetStateReq) })
	c.regMsg(DigActivityChangeGadgetStateRsp, func() any { return new(proto.DigActivityChangeGadgetStateRsp) })
	c.regMsg(DigActivityMarkPointChangeNotify, func() any { return new(proto.DigActivityMarkPointChangeNotify) })
	c.regMsg(DisableRoguelikeTrapNotify, func() any { return new(proto.DisableRoguelikeTrapNotify) })
	c.regMsg(DoGachaReq, func() any { return new(proto.DoGachaReq) })
	c.regMsg(DoGachaRsp, func() any { return new(proto.DoGachaRsp) })
	c.regMsg(DoRoguelikeDungeonCardGachaReq, func() any { return new(proto.DoRoguelikeDungeonCardGachaReq) })
	c.regMsg(DoRoguelikeDungeonCardGachaRsp, func() any { return new(proto.DoRoguelikeDungeonCardGachaRsp) })
	c.regMsg(DoSetPlayerBornDataNotify, func() any { return new(proto.DoSetPlayerBornDataNotify) })
	c.regMsg(DraftGuestReplyInviteNotify, func() any { return new(proto.DraftGuestReplyInviteNotify) })
	c.regMsg(DraftGuestReplyInviteReq, func() any { return new(proto.DraftGuestReplyInviteReq) })
	c.regMsg(DraftGuestReplyInviteRsp, func() any { return new(proto.DraftGuestReplyInviteRsp) })
	c.regMsg(DraftGuestReplyTwiceConfirmNotify, func() any { return new(proto.DraftGuestReplyTwiceConfirmNotify) })
	c.regMsg(DraftGuestReplyTwiceConfirmReq, func() any { return new(proto.DraftGuestReplyTwiceConfirmReq) })
	c.regMsg(DraftGuestReplyTwiceConfirmRsp, func() any { return new(proto.DraftGuestReplyTwiceConfirmRsp) })
	c.regMsg(DraftInviteResultNotify, func() any { return new(proto.DraftInviteResultNotify) })
	c.regMsg(DraftOwnerInviteNotify, func() any { return new(proto.DraftOwnerInviteNotify) })
	c.regMsg(DraftOwnerStartInviteReq, func() any { return new(proto.DraftOwnerStartInviteReq) })
	c.regMsg(DraftOwnerStartInviteRsp, func() any { return new(proto.DraftOwnerStartInviteRsp) })
	c.regMsg(DraftOwnerTwiceConfirmNotify, func() any { return new(proto.DraftOwnerTwiceConfirmNotify) })
	c.regMsg(DraftTwiceConfirmResultNotify, func() any { return new(proto.DraftTwiceConfirmResultNotify) })
	c.regMsg(DragonSpineChapterFinishNotify, func() any { return new(proto.DragonSpineChapterFinishNotify) })
	c.regMsg(DragonSpineChapterOpenNotify, func() any { return new(proto.DragonSpineChapterOpenNotify) })
	c.regMsg(DragonSpineChapterProgressChangeNotify, func() any { return new(proto.DragonSpineChapterProgressChangeNotify) })
	c.regMsg(DragonSpineCoinChangeNotify, func() any { return new(proto.DragonSpineCoinChangeNotify) })
	c.regMsg(DropHintNotify, func() any { return new(proto.DropHintNotify) })
	c.regMsg(DropItemReq, func() any { return new(proto.DropItemReq) })
	c.regMsg(DropItemRsp, func() any { return new(proto.DropItemRsp) })
	c.regMsg(DungeonCandidateTeamChangeAvatarReq, func() any { return new(proto.DungeonCandidateTeamChangeAvatarReq) })
	c.regMsg(DungeonCandidateTeamChangeAvatarRsp, func() any { return new(proto.DungeonCandidateTeamChangeAvatarRsp) })
	c.regMsg(DungeonCandidateTeamCreateReq, func() any { return new(proto.DungeonCandidateTeamCreateReq) })
	c.regMsg(DungeonCandidateTeamCreateRsp, func() any { return new(proto.DungeonCandidateTeamCreateRsp) })
	c.regMsg(DungeonCandidateTeamDismissNotify, func() any { return new(proto.DungeonCandidateTeamDismissNotify) })
	c.regMsg(DungeonCandidateTeamInfoNotify, func() any { return new(proto.DungeonCandidateTeamInfoNotify) })
	c.regMsg(DungeonCandidateTeamInviteNotify, func() any { return new(proto.DungeonCandidateTeamInviteNotify) })
	c.regMsg(DungeonCandidateTeamInviteReq, func() any { return new(proto.DungeonCandidateTeamInviteReq) })
	c.regMsg(DungeonCandidateTeamInviteRsp, func() any { return new(proto.DungeonCandidateTeamInviteRsp) })
	c.regMsg(DungeonCandidateTeamKickReq, func() any { return new(proto.DungeonCandidateTeamKickReq) })
	c.regMsg(DungeonCandidateTeamKickRsp, func() any { return new(proto.DungeonCandidateTeamKickRsp) })
	c.regMsg(DungeonCandidateTeamLeaveReq, func() any { return new(proto.DungeonCandidateTeamLeaveReq) })
	c.regMsg(DungeonCandidateTeamLeaveRsp, func() any { return new(proto.DungeonCandidateTeamLeaveRsp) })
	c.regMsg(DungeonCandidateTeamPlayerLeaveNotify, func() any { return new(proto.DungeonCandidateTeamPlayerLeaveNotify) })
	c.regMsg(DungeonCandidateTeamRefuseNotify, func() any { return new(proto.DungeonCandidateTeamRefuseNotify) })
	c.regMsg(DungeonCandidateTeamReplyInviteReq, func() any { return new(proto.DungeonCandidateTeamReplyInviteReq) })
	c.regMsg(DungeonCandidateTeamReplyInviteRsp, func() any { return new(proto.DungeonCandidateTeamReplyInviteRsp) })
	c.regMsg(DungeonCandidateTeamSetChangingAvatarReq, func() any { return new(proto.DungeonCandidateTeamSetChangingAvatarReq) })
	c.regMsg(DungeonCandidateTeamSetChangingAvatarRsp, func() any { return new(proto.DungeonCandidateTeamSetChangingAvatarRsp) })
	c.regMsg(DungeonCandidateTeamSetReadyReq, func() any { return new(proto.DungeonCandidateTeamSetReadyReq) })
	c.regMsg(DungeonCandidateTeamSetReadyRsp, func() any { return new(proto.DungeonCandidateTeamSetReadyRsp) })
	c.regMsg(DungeonChallengeBeginNotify, func() any { return new(proto.DungeonChallengeBeginNotify) })
	c.regMsg(DungeonChallengeFinishNotify, func() any { return new(proto.DungeonChallengeFinishNotify) })
	c.regMsg(DungeonDataNotify, func() any { return new(proto.DungeonDataNotify) })
	c.regMsg(DungeonDieOptionReq, func() any { return new(proto.DungeonDieOptionReq) })
	c.regMsg(DungeonDieOptionRsp, func() any { return new(proto.DungeonDieOptionRsp) })
	c.regMsg(DungeonEntryInfoReq, func() any { return new(proto.DungeonEntryInfoReq) })
	c.regMsg(DungeonEntryInfoRsp, func() any { return new(proto.DungeonEntryInfoRsp) })
	c.regMsg(DungeonEntryToBeExploreNotify, func() any { return new(proto.DungeonEntryToBeExploreNotify) })
	c.regMsg(DungeonFollowNotify, func() any { return new(proto.DungeonFollowNotify) })
	c.regMsg(DungeonGetStatueDropReq, func() any { return new(proto.DungeonGetStatueDropReq) })
	c.regMsg(DungeonGetStatueDropRsp, func() any { return new(proto.DungeonGetStatueDropRsp) })
	c.regMsg(DungeonInterruptChallengeReq, func() any { return new(proto.DungeonInterruptChallengeReq) })
	c.regMsg(DungeonInterruptChallengeRsp, func() any { return new(proto.DungeonInterruptChallengeRsp) })
	c.regMsg(DungeonPlayerDieNotify, func() any { return new(proto.DungeonPlayerDieNotify) })
	c.regMsg(DungeonPlayerDieReq, func() any { return new(proto.DungeonPlayerDieReq) })
	c.regMsg(DungeonPlayerDieRsp, func() any { return new(proto.DungeonPlayerDieRsp) })
	c.regMsg(DungeonRestartInviteNotify, func() any { return new(proto.DungeonRestartInviteNotify) })
	c.regMsg(DungeonRestartInviteReplyNotify, func() any { return new(proto.DungeonRestartInviteReplyNotify) })
	c.regMsg(DungeonRestartInviteReplyReq, func() any { return new(proto.DungeonRestartInviteReplyReq) })
	c.regMsg(DungeonRestartInviteReplyRsp, func() any { return new(proto.DungeonRestartInviteReplyRsp) })
	c.regMsg(DungeonRestartReq, func() any { return new(proto.DungeonRestartReq) })
	c.regMsg(DungeonRestartResultNotify, func() any { return new(proto.DungeonRestartResultNotify) })
	c.regMsg(DungeonRestartRsp, func() any { return new(proto.DungeonRestartRsp) })
	c.regMsg(DungeonReviseLevelNotify, func() any { return new(proto.DungeonReviseLevelNotify) })
	c.regMsg(DungeonSettleNotify, func() any { return new(proto.DungeonSettleNotify) })
	c.regMsg(DungeonShowReminderNotify, func() any { return new(proto.DungeonShowReminderNotify) })
	c.regMsg(DungeonSlipRevivePointActivateReq, func() any { return new(proto.DungeonSlipRevivePointActivateReq) })
	c.regMsg(DungeonSlipRevivePointActivateRsp, func() any { return new(proto.DungeonSlipRevivePointActivateRsp) })
	c.regMsg(DungeonWayPointActivateReq, func() any { return new(proto.DungeonWayPointActivateReq) })
	c.regMsg(DungeonWayPointActivateRsp, func() any { return new(proto.DungeonWayPointActivateRsp) })
	c.regMsg(DungeonWayPointNotify, func() any { return new(proto.DungeonWayPointNotify) })
	c.regMsg(EchoNotify, func() any { return new(proto.EchoNotify) })
	c.regMsg(EchoShellTakeRewardReq, func() any { return new(proto.EchoShellTakeRewardReq) })
	c.regMsg(EchoShellTakeRewardRsp, func() any { return new(proto.EchoShellTakeRewardRsp) })
	c.regMsg(EchoShellUpdateNotify, func() any { return new(proto.EchoShellUpdateNotify) })
	c.regMsg(EffigyChallengeInfoNotify, func() any { return new(proto.EffigyChallengeInfoNotify) })
	c.regMsg(EffigyChallengeResultNotify, func() any { return new(proto.EffigyChallengeResultNotify) })
	c.regMsg(EffigyChallengeV2ChooseSkillReq, func() any { return new(proto.EffigyChallengeV2ChooseSkillReq) })
	c.regMsg(EffigyChallengeV2ChooseSkillRsp, func() any { return new(proto.EffigyChallengeV2ChooseSkillRsp) })
	c.regMsg(EffigyChallengeV2DungeonInfoNotify, func() any { return new(proto.EffigyChallengeV2DungeonInfoNotify) })
	c.regMsg(EffigyChallengeV2EnterDungeonReq, func() any { return new(proto.EffigyChallengeV2EnterDungeonReq) })
	c.regMsg(EffigyChallengeV2EnterDungeonRsp, func() any { return new(proto.EffigyChallengeV2EnterDungeonRsp) })
	c.regMsg(EffigyChallengeV2RestartDungeonReq, func() any { return new(proto.EffigyChallengeV2RestartDungeonReq) })
	c.regMsg(EffigyChallengeV2RestartDungeonRsp, func() any { return new(proto.EffigyChallengeV2RestartDungeonRsp) })
	c.regMsg(EndCameraSceneLookNotify, func() any { return new(proto.EndCameraSceneLookNotify) })
	c.regMsg(EnterChessDungeonReq, func() any { return new(proto.EnterChessDungeonReq) })
	c.regMsg(EnterChessDungeonRsp, func() any { return new(proto.EnterChessDungeonRsp) })
	c.regMsg(EnterCustomDungeonReq, func() any { return new(proto.EnterCustomDungeonReq) })
	c.regMsg(EnterCustomDungeonRsp, func() any { return new(proto.EnterCustomDungeonRsp) })
	c.regMsg(EnterFishingReq, func() any { return new(proto.EnterFishingReq) })
	c.regMsg(EnterFishingRsp, func() any { return new(proto.EnterFishingRsp) })
	c.regMsg(EnterFungusFighterPlotDungeonReq, func() any { return new(proto.EnterFungusFighterPlotDungeonReq) })
	c.regMsg(EnterFungusFighterPlotDungeonRsp, func() any { return new(proto.EnterFungusFighterPlotDungeonRsp) })
	c.regMsg(EnterFungusFighterTrainingDungeonReq, func() any { return new(proto.EnterFungusFighterTrainingDungeonReq) })
	c.regMsg(EnterFungusFighterTrainingDungeonRsp, func() any { return new(proto.EnterFungusFighterTrainingDungeonRsp) })
	c.regMsg(EnterIrodoriChessDungeonReq, func() any { return new(proto.EnterIrodoriChessDungeonReq) })
	c.regMsg(EnterIrodoriChessDungeonRsp, func() any { return new(proto.EnterIrodoriChessDungeonRsp) })
	c.regMsg(EnterMechanicusDungeonReq, func() any { return new(proto.EnterMechanicusDungeonReq) })
	c.regMsg(EnterMechanicusDungeonRsp, func() any { return new(proto.EnterMechanicusDungeonRsp) })
	c.regMsg(EnterRogueDiaryDungeonReq, func() any { return new(proto.EnterRogueDiaryDungeonReq) })
	c.regMsg(EnterRogueDiaryDungeonRsp, func() any { return new(proto.EnterRogueDiaryDungeonRsp) })
	c.regMsg(EnterRoguelikeDungeonNotify, func() any { return new(proto.EnterRoguelikeDungeonNotify) })
	c.regMsg(EnterSceneDoneReq, func() any { return new(proto.EnterSceneDoneReq) })
	c.regMsg(EnterSceneDoneRsp, func() any { return new(proto.EnterSceneDoneRsp) })
	c.regMsg(EnterScenePeerNotify, func() any { return new(proto.EnterScenePeerNotify) })
	c.regMsg(EnterSceneReadyReq, func() any { return new(proto.EnterSceneReadyReq) })
	c.regMsg(EnterSceneReadyRsp, func() any { return new(proto.EnterSceneReadyRsp) })
	c.regMsg(EnterSceneWeatherAreaNotify, func() any { return new(proto.EnterSceneWeatherAreaNotify) })
	c.regMsg(EnterTransPointRegionNotify, func() any { return new(proto.EnterTransPointRegionNotify) })
	c.regMsg(EnterTrialAvatarActivityDungeonReq, func() any { return new(proto.EnterTrialAvatarActivityDungeonReq) })
	c.regMsg(EnterTrialAvatarActivityDungeonRsp, func() any { return new(proto.EnterTrialAvatarActivityDungeonRsp) })
	c.regMsg(EnterWorldAreaReq, func() any { return new(proto.EnterWorldAreaReq) })
	c.regMsg(EnterWorldAreaRsp, func() any { return new(proto.EnterWorldAreaRsp) })
	c.regMsg(EntityAiKillSelfNotify, func() any { return new(proto.EntityAiKillSelfNotify) })
	c.regMsg(EntityAiSyncNotify, func() any { return new(proto.EntityAiSyncNotify) })
	c.regMsg(EntityAuthorityChangeNotify, func() any { return new(proto.EntityAuthorityChangeNotify) })
	c.regMsg(EntityConfigHashNotify, func() any { return new(proto.EntityConfigHashNotify) })
	c.regMsg(EntityFightPropChangeReasonNotify, func() any { return new(proto.EntityFightPropChangeReasonNotify) })
	c.regMsg(EntityFightPropNotify, func() any { return new(proto.EntityFightPropNotify) })
	c.regMsg(EntityFightPropUpdateNotify, func() any { return new(proto.EntityFightPropUpdateNotify) })
	c.regMsg(EntityForceSyncReq, func() any { return new(proto.EntityForceSyncReq) })
	c.regMsg(EntityForceSyncRsp, func() any { return new(proto.EntityForceSyncRsp) })
	c.regMsg(EntityJumpNotify, func() any { return new(proto.EntityJumpNotify) })
	c.regMsg(EntityMoveRoomNotify, func() any { return new(proto.EntityMoveRoomNotify) })
	c.regMsg(EntityPropNotify, func() any { return new(proto.EntityPropNotify) })
	c.regMsg(EntityTagChangeNotify, func() any { return new(proto.EntityTagChangeNotify) })
	c.regMsg(EquipRoguelikeRuneReq, func() any { return new(proto.EquipRoguelikeRuneReq) })
	c.regMsg(EquipRoguelikeRuneRsp, func() any { return new(proto.EquipRoguelikeRuneRsp) })
	c.regMsg(EvtAiSyncCombatThreatInfoNotify, func() any { return new(proto.EvtAiSyncCombatThreatInfoNotify) })
	c.regMsg(EvtAiSyncSkillCdNotify, func() any { return new(proto.EvtAiSyncSkillCdNotify) })
	c.regMsg(EvtAnimatorParameterNotify, func() any { return new(proto.EvtAnimatorParameterNotify) })
	c.regMsg(EvtAnimatorStateChangedNotify, func() any { return new(proto.EvtAnimatorStateChangedNotify) })
	c.regMsg(EvtAvatarEnterFocusNotify, func() any { return new(proto.EvtAvatarEnterFocusNotify) })
	c.regMsg(EvtAvatarExitFocusNotify, func() any { return new(proto.EvtAvatarExitFocusNotify) })
	c.regMsg(EvtAvatarLockChairReq, func() any { return new(proto.EvtAvatarLockChairReq) })
	c.regMsg(EvtAvatarLockChairRsp, func() any { return new(proto.EvtAvatarLockChairRsp) })
	c.regMsg(EvtAvatarSitDownNotify, func() any { return new(proto.EvtAvatarSitDownNotify) })
	c.regMsg(EvtAvatarStandUpNotify, func() any { return new(proto.EvtAvatarStandUpNotify) })
	c.regMsg(EvtAvatarUpdateFocusNotify, func() any { return new(proto.EvtAvatarUpdateFocusNotify) })
	c.regMsg(EvtBeingHealedNotify, func() any { return new(proto.EvtBeingHealedNotify) })
	c.regMsg(EvtBeingHitNotify, func() any { return new(proto.EvtBeingHitNotify) })
	c.regMsg(EvtBeingHitsCombineNotify, func() any { return new(proto.EvtBeingHitsCombineNotify) })
	c.regMsg(EvtBulletDeactiveNotify, func() any { return new(proto.EvtBulletDeactiveNotify) })
	c.regMsg(EvtBulletHitNotify, func() any { return new(proto.EvtBulletHitNotify) })
	c.regMsg(EvtBulletMoveNotify, func() any { return new(proto.EvtBulletMoveNotify) })
	c.regMsg(EvtCostStaminaNotify, func() any { return new(proto.EvtCostStaminaNotify) })
	c.regMsg(EvtCreateGadgetNotify, func() any { return new(proto.EvtCreateGadgetNotify) })
	c.regMsg(EvtDestroyGadgetNotify, func() any { return new(proto.EvtDestroyGadgetNotify) })
	c.regMsg(EvtDestroyServerGadgetNotify, func() any { return new(proto.EvtDestroyServerGadgetNotify) })
	c.regMsg(EvtDoSkillSuccNotify, func() any { return new(proto.EvtDoSkillSuccNotify) })
	c.regMsg(EvtEntityRenderersChangedNotify, func() any { return new(proto.EvtEntityRenderersChangedNotify) })
	c.regMsg(EvtEntityStartDieEndNotify, func() any { return new(proto.EvtEntityStartDieEndNotify) })
	c.regMsg(EvtFaceToDirNotify, func() any { return new(proto.EvtFaceToDirNotify) })
	c.regMsg(EvtFaceToEntityNotify, func() any { return new(proto.EvtFaceToEntityNotify) })
	c.regMsg(EvtLocalGadgetOwnerLeaveSceneNotify, func() any { return new(proto.EvtLocalGadgetOwnerLeaveSceneNotify) })
	c.regMsg(EvtRushMoveNotify, func() any { return new(proto.EvtRushMoveNotify) })
	c.regMsg(EvtSetAttackTargetNotify, func() any { return new(proto.EvtSetAttackTargetNotify) })
	c.regMsg(ExclusiveRuleNotify, func() any { return new(proto.ExclusiveRuleNotify) })
	c.regMsg(ExecuteGadgetLuaReq, func() any { return new(proto.ExecuteGadgetLuaReq) })
	c.regMsg(ExecuteGadgetLuaRsp, func() any { return new(proto.ExecuteGadgetLuaRsp) })
	c.regMsg(ExecuteGroupTriggerReq, func() any { return new(proto.ExecuteGroupTriggerReq) })
	c.regMsg(ExecuteGroupTriggerRsp, func() any { return new(proto.ExecuteGroupTriggerRsp) })
	c.regMsg(ExitCustomDungeonTryReq, func() any { return new(proto.ExitCustomDungeonTryReq) })
	c.regMsg(ExitCustomDungeonTryRsp, func() any { return new(proto.ExitCustomDungeonTryRsp) })
	c.regMsg(ExitFishingReq, func() any { return new(proto.ExitFishingReq) })
	c.regMsg(ExitFishingRsp, func() any { return new(proto.ExitFishingRsp) })
	c.regMsg(ExitSceneWeatherAreaNotify, func() any { return new(proto.ExitSceneWeatherAreaNotify) })
	c.regMsg(ExitTransPointRegionNotify, func() any { return new(proto.ExitTransPointRegionNotify) })
	c.regMsg(ExpeditionChallengeEnterRegionNotify, func() any { return new(proto.ExpeditionChallengeEnterRegionNotify) })
	c.regMsg(ExpeditionChallengeFinishedNotify, func() any { return new(proto.ExpeditionChallengeFinishedNotify) })
	c.regMsg(ExpeditionRecallReq, func() any { return new(proto.ExpeditionRecallReq) })
	c.regMsg(ExpeditionRecallRsp, func() any { return new(proto.ExpeditionRecallRsp) })
	c.regMsg(ExpeditionStartReq, func() any { return new(proto.ExpeditionStartReq) })
	c.regMsg(ExpeditionStartRsp, func() any { return new(proto.ExpeditionStartRsp) })
	c.regMsg(ExpeditionTakeRewardReq, func() any { return new(proto.ExpeditionTakeRewardReq) })
	c.regMsg(ExpeditionTakeRewardRsp, func() any { return new(proto.ExpeditionTakeRewardRsp) })
	c.regMsg(FindHilichurlAcceptQuestNotify, func() any { return new(proto.FindHilichurlAcceptQuestNotify) })
	c.regMsg(FindHilichurlFinishSecondQuestNotify, func() any { return new(proto.FindHilichurlFinishSecondQuestNotify) })
	c.regMsg(FinishDeliveryNotify, func() any { return new(proto.FinishDeliveryNotify) })
	c.regMsg(FinishLanternProjectionReq, func() any { return new(proto.FinishLanternProjectionReq) })
	c.regMsg(FinishLanternProjectionRsp, func() any { return new(proto.FinishLanternProjectionRsp) })
	c.regMsg(FinishMainCoopReq, func() any { return new(proto.FinishMainCoopReq) })
	c.regMsg(FinishMainCoopRsp, func() any { return new(proto.FinishMainCoopRsp) })
	c.regMsg(FinishedParentQuestNotify, func() any { return new(proto.FinishedParentQuestNotify) })
	c.regMsg(FinishedParentQuestUpdateNotify, func() any { return new(proto.FinishedParentQuestUpdateNotify) })
	c.regMsg(FinishedTalkIdListNotify, func() any { return new(proto.FinishedTalkIdListNotify) })
	c.regMsg(FireworksLaunchDataNotify, func() any { return new(proto.FireworksLaunchDataNotify) })
	c.regMsg(FireworksReformDataNotify, func() any { return new(proto.FireworksReformDataNotify) })
	c.regMsg(FishAttractNotify, func() any { return new(proto.FishAttractNotify) })
	c.regMsg(FishBaitGoneNotify, func() any { return new(proto.FishBaitGoneNotify) })
	c.regMsg(FishBattleBeginReq, func() any { return new(proto.FishBattleBeginReq) })
	c.regMsg(FishBattleBeginRsp, func() any { return new(proto.FishBattleBeginRsp) })
	c.regMsg(FishBattleEndReq, func() any { return new(proto.FishBattleEndReq) })
	c.regMsg(FishBattleEndRsp, func() any { return new(proto.FishBattleEndRsp) })
	c.regMsg(FishBiteReq, func() any { return new(proto.FishBiteReq) })
	c.regMsg(FishBiteRsp, func() any { return new(proto.FishBiteRsp) })
	c.regMsg(FishCastRodReq, func() any { return new(proto.FishCastRodReq) })
	c.regMsg(FishCastRodRsp, func() any { return new(proto.FishCastRodRsp) })
	c.regMsg(FishChosenNotify, func() any { return new(proto.FishChosenNotify) })
	c.regMsg(FishEscapeNotify, func() any { return new(proto.FishEscapeNotify) })
	c.regMsg(FishPoolDataNotify, func() any { return new(proto.FishPoolDataNotify) })
	c.regMsg(FishingGallerySettleNotify, func() any { return new(proto.FishingGallerySettleNotify) })
	c.regMsg(FleurFairBalloonSettleNotify, func() any { return new(proto.FleurFairBalloonSettleNotify) })
	c.regMsg(FleurFairBuffEnergyNotify, func() any { return new(proto.FleurFairBuffEnergyNotify) })
	c.regMsg(FleurFairFallSettleNotify, func() any { return new(proto.FleurFairFallSettleNotify) })
	c.regMsg(FleurFairFinishGalleryStageNotify, func() any { return new(proto.FleurFairFinishGalleryStageNotify) })
	c.regMsg(FleurFairMusicGameSettleReq, func() any { return new(proto.FleurFairMusicGameSettleReq) })
	c.regMsg(FleurFairMusicGameSettleRsp, func() any { return new(proto.FleurFairMusicGameSettleRsp) })
	c.regMsg(FleurFairMusicGameStartReq, func() any { return new(proto.FleurFairMusicGameStartReq) })
	c.regMsg(FleurFairMusicGameStartRsp, func() any { return new(proto.FleurFairMusicGameStartRsp) })
	c.regMsg(FleurFairReplayMiniGameReq, func() any { return new(proto.FleurFairReplayMiniGameReq) })
	c.regMsg(FleurFairReplayMiniGameRsp, func() any { return new(proto.FleurFairReplayMiniGameRsp) })
	c.regMsg(FleurFairStageSettleNotify, func() any { return new(proto.FleurFairStageSettleNotify) })
	c.regMsg(FlightActivityRestartReq, func() any { return new(proto.FlightActivityRestartReq) })
	c.regMsg(FlightActivityRestartRsp, func() any { return new(proto.FlightActivityRestartRsp) })
	c.regMsg(FlightActivitySettleNotify, func() any { return new(proto.FlightActivitySettleNotify) })
	c.regMsg(FocusAvatarReq, func() any { return new(proto.FocusAvatarReq) })
	c.regMsg(FocusAvatarRsp, func() any { return new(proto.FocusAvatarRsp) })
	c.regMsg(ForceAddPlayerFriendReq, func() any { return new(proto.ForceAddPlayerFriendReq) })
	c.regMsg(ForceAddPlayerFriendRsp, func() any { return new(proto.ForceAddPlayerFriendRsp) })
	c.regMsg(ForceDragAvatarNotify, func() any { return new(proto.ForceDragAvatarNotify) })
	c.regMsg(ForceDragBackTransferNotify, func() any { return new(proto.ForceDragBackTransferNotify) })
	c.regMsg(ForgeDataNotify, func() any { return new(proto.ForgeDataNotify) })
	c.regMsg(ForgeFormulaDataNotify, func() any { return new(proto.ForgeFormulaDataNotify) })
	c.regMsg(ForgeGetQueueDataReq, func() any { return new(proto.ForgeGetQueueDataReq) })
	c.regMsg(ForgeGetQueueDataRsp, func() any { return new(proto.ForgeGetQueueDataRsp) })
	c.regMsg(ForgeQueueDataNotify, func() any { return new(proto.ForgeQueueDataNotify) })
	c.regMsg(ForgeQueueManipulateReq, func() any { return new(proto.ForgeQueueManipulateReq) })
	c.regMsg(ForgeQueueManipulateRsp, func() any { return new(proto.ForgeQueueManipulateRsp) })
	c.regMsg(ForgeStartReq, func() any { return new(proto.ForgeStartReq) })
	c.regMsg(ForgeStartRsp, func() any { return new(proto.ForgeStartRsp) })
	c.regMsg(FoundationNotify, func() any { return new(proto.FoundationNotify) })
	c.regMsg(FoundationReq, func() any { return new(proto.FoundationReq) })
	c.regMsg(FoundationRsp, func() any { return new(proto.FoundationRsp) })
	c.regMsg(FriendInfoChangeNotify, func() any { return new(proto.FriendInfoChangeNotify) })
	c.regMsg(FungusCaptureSettleNotify, func() any { return new(proto.FungusCaptureSettleNotify) })
	c.regMsg(FungusCultivateReq, func() any { return new(proto.FungusCultivateReq) })
	c.regMsg(FungusCultivateRsp, func() any { return new(proto.FungusCultivateRsp) })
	c.regMsg(FungusFighterClearTrainingRuntimeDataReq, func() any { return new(proto.FungusFighterClearTrainingRuntimeDataReq) })
	c.regMsg(FungusFighterClearTrainingRuntimeDataRsp, func() any { return new(proto.FungusFighterClearTrainingRuntimeDataRsp) })
	c.regMsg(FungusFighterPlotInfoNotify, func() any { return new(proto.FungusFighterPlotInfoNotify) })
	c.regMsg(FungusFighterRuntimeDataNotify, func() any { return new(proto.FungusFighterRuntimeDataNotify) })
	c.regMsg(FungusFighterTrainingGallerySettleNotify, func() any { return new(proto.FungusFighterTrainingGallerySettleNotify) })
	c.regMsg(FungusFighterTrainingInfoNotify, func() any { return new(proto.FungusFighterTrainingInfoNotify) })
	c.regMsg(FungusFighterTrainingSelectFungusReq, func() any { return new(proto.FungusFighterTrainingSelectFungusReq) })
	c.regMsg(FungusFighterTrainingSelectFungusRsp, func() any { return new(proto.FungusFighterTrainingSelectFungusRsp) })
	c.regMsg(FungusFighterUseBackupFungusReq, func() any { return new(proto.FungusFighterUseBackupFungusReq) })
	c.regMsg(FungusFighterUseBackupFungusRsp, func() any { return new(proto.FungusFighterUseBackupFungusRsp) })
	c.regMsg(FungusRenameReq, func() any { return new(proto.FungusRenameReq) })
	c.regMsg(FungusRenameRsp, func() any { return new(proto.FungusRenameRsp) })
	c.regMsg(FurnitureCurModuleArrangeCountNotify, func() any { return new(proto.FurnitureCurModuleArrangeCountNotify) })
	c.regMsg(FurnitureMakeBeHelpedNotify, func() any { return new(proto.FurnitureMakeBeHelpedNotify) })
	c.regMsg(FurnitureMakeCancelReq, func() any { return new(proto.FurnitureMakeCancelReq) })
	c.regMsg(FurnitureMakeCancelRsp, func() any { return new(proto.FurnitureMakeCancelRsp) })
	c.regMsg(FurnitureMakeFinishNotify, func() any { return new(proto.FurnitureMakeFinishNotify) })
	c.regMsg(FurnitureMakeHelpReq, func() any { return new(proto.FurnitureMakeHelpReq) })
	c.regMsg(FurnitureMakeHelpRsp, func() any { return new(proto.FurnitureMakeHelpRsp) })
	c.regMsg(FurnitureMakeReq, func() any { return new(proto.FurnitureMakeReq) })
	c.regMsg(FurnitureMakeRsp, func() any { return new(proto.FurnitureMakeRsp) })
	c.regMsg(FurnitureMakeStartReq, func() any { return new(proto.FurnitureMakeStartReq) })
	c.regMsg(FurnitureMakeStartRsp, func() any { return new(proto.FurnitureMakeStartRsp) })
	c.regMsg(GCGApplyInviteBattleNotify, func() any { return new(proto.GCGApplyInviteBattleNotify) })
	c.regMsg(GCGApplyInviteBattleReq, func() any { return new(proto.GCGApplyInviteBattleReq) })
	c.regMsg(GCGApplyInviteBattleRsp, func() any { return new(proto.GCGApplyInviteBattleRsp) })
	c.regMsg(GCGAskDuelReq, func() any { return new(proto.GCGAskDuelReq) })
	c.regMsg(GCGAskDuelRsp, func() any { return new(proto.GCGAskDuelRsp) })
	c.regMsg(GCGBasicDataNotify, func() any { return new(proto.GCGBasicDataNotify) })
	c.regMsg(GCGBossChallengeUpdateNotify, func() any { return new(proto.GCGBossChallengeUpdateNotify) })
	c.regMsg(GCGChallengeUpdateNotify, func() any { return new(proto.GCGChallengeUpdateNotify) })
	c.regMsg(GCGClientSettleReq, func() any { return new(proto.GCGClientSettleReq) })
	c.regMsg(GCGClientSettleRsp, func() any { return new(proto.GCGClientSettleRsp) })
	c.regMsg(GCGDSCardBackUnlockNotify, func() any { return new(proto.GCGDSCardBackUnlockNotify) })
	c.regMsg(GCGDSCardFaceUnlockNotify, func() any { return new(proto.GCGDSCardFaceUnlockNotify) })
	c.regMsg(GCGDSCardNumChangeNotify, func() any { return new(proto.GCGDSCardNumChangeNotify) })
	c.regMsg(GCGDSCardProficiencyNotify, func() any { return new(proto.GCGDSCardProficiencyNotify) })
	c.regMsg(GCGDSChangeCardBackReq, func() any { return new(proto.GCGDSChangeCardBackReq) })
	c.regMsg(GCGDSChangeCardBackRsp, func() any { return new(proto.GCGDSChangeCardBackRsp) })
	c.regMsg(GCGDSChangeCardFaceReq, func() any { return new(proto.GCGDSChangeCardFaceReq) })
	c.regMsg(GCGDSChangeCardFaceRsp, func() any { return new(proto.GCGDSChangeCardFaceRsp) })
	c.regMsg(GCGDSChangeCurDeckReq, func() any { return new(proto.GCGDSChangeCurDeckReq) })
	c.regMsg(GCGDSChangeCurDeckRsp, func() any { return new(proto.GCGDSChangeCurDeckRsp) })
	c.regMsg(GCGDSChangeDeckNameReq, func() any { return new(proto.GCGDSChangeDeckNameReq) })
	c.regMsg(GCGDSChangeDeckNameRsp, func() any { return new(proto.GCGDSChangeDeckNameRsp) })
	c.regMsg(GCGDSChangeFieldReq, func() any { return new(proto.GCGDSChangeFieldReq) })
	c.regMsg(GCGDSChangeFieldRsp, func() any { return new(proto.GCGDSChangeFieldRsp) })
	c.regMsg(GCGDSCurDeckChangeNotify, func() any { return new(proto.GCGDSCurDeckChangeNotify) })
	c.regMsg(GCGDSDataNotify, func() any { return new(proto.GCGDSDataNotify) })
	c.regMsg(GCGDSDeckSaveReq, func() any { return new(proto.GCGDSDeckSaveReq) })
	c.regMsg(GCGDSDeckSaveRsp, func() any { return new(proto.GCGDSDeckSaveRsp) })
	c.regMsg(GCGDSDeckUnlockNotify, func() any { return new(proto.GCGDSDeckUnlockNotify) })
	c.regMsg(GCGDSDeleteDeckReq, func() any { return new(proto.GCGDSDeleteDeckReq) })
	c.regMsg(GCGDSDeleteDeckRsp, func() any { return new(proto.GCGDSDeleteDeckRsp) })
	c.regMsg(GCGDSFieldUnlockNotify, func() any { return new(proto.GCGDSFieldUnlockNotify) })
	c.regMsg(GCGGameBriefDataNotify, func() any { return new(proto.GCGGameBriefDataNotify) })
	c.regMsg(GCGGrowthLevelNotify, func() any { return new(proto.GCGGrowthLevelNotify) })
	c.regMsg(GCGGrowthLevelRewardNotify, func() any { return new(proto.GCGGrowthLevelRewardNotify) })
	c.regMsg(GCGGrowthLevelTakeRewardReq, func() any { return new(proto.GCGGrowthLevelTakeRewardReq) })
	c.regMsg(GCGGrowthLevelTakeRewardRsp, func() any { return new(proto.GCGGrowthLevelTakeRewardRsp) })
	c.regMsg(GCGHeartBeatNotify, func() any { return new(proto.GCGHeartBeatNotify) })
	c.regMsg(GCGInitFinishReq, func() any { return new(proto.GCGInitFinishReq) })
	c.regMsg(GCGInitFinishRsp, func() any { return new(proto.GCGInitFinishRsp) })
	c.regMsg(GCGInviteBattleNotify, func() any { return new(proto.GCGInviteBattleNotify) })
	c.regMsg(GCGInviteGuestBattleReq, func() any { return new(proto.GCGInviteGuestBattleReq) })
	c.regMsg(GCGInviteGuestBattleRsp, func() any { return new(proto.GCGInviteGuestBattleRsp) })
	c.regMsg(GCGLevelChallengeFinishNotify, func() any { return new(proto.GCGLevelChallengeFinishNotify) })
	c.regMsg(GCGLevelChallengeNotify, func() any { return new(proto.GCGLevelChallengeNotify) })
	c.regMsg(GCGMessagePackNotify, func() any { return new(proto.GCGMessagePackNotify) })
	c.regMsg(GCGOperationReq, func() any { return new(proto.GCGOperationReq) })
	c.regMsg(GCGOperationRsp, func() any { return new(proto.GCGOperationRsp) })
	c.regMsg(GCGResourceStateNotify, func() any { return new(proto.GCGResourceStateNotify) })
	c.regMsg(GCGSettleNotify, func() any { return new(proto.GCGSettleNotify) })
	c.regMsg(GCGSettleOptionReq, func() any { return new(proto.GCGSettleOptionReq) })
	c.regMsg(GCGSettleOptionRsp, func() any { return new(proto.GCGSettleOptionRsp) })
	c.regMsg(GCGSkillPreviewAskReq, func() any { return new(proto.GCGSkillPreviewAskReq) })
	c.regMsg(GCGSkillPreviewAskRsp, func() any { return new(proto.GCGSkillPreviewAskRsp) })
	c.regMsg(GCGSkillPreviewNotify, func() any { return new(proto.GCGSkillPreviewNotify) })
	c.regMsg(GCGStartChallengeReq, func() any { return new(proto.GCGStartChallengeReq) })
	c.regMsg(GCGStartChallengeRsp, func() any { return new(proto.GCGStartChallengeRsp) })
	c.regMsg(GCGTCInviteReq, func() any { return new(proto.GCGTCInviteReq) })
	c.regMsg(GCGTCInviteRsp, func() any { return new(proto.GCGTCInviteRsp) })
	c.regMsg(GCGTCTavernChallengeDataNotify, func() any { return new(proto.GCGTCTavernChallengeDataNotify) })
	c.regMsg(GCGTCTavernChallengeUpdateNotify, func() any { return new(proto.GCGTCTavernChallengeUpdateNotify) })
	c.regMsg(GCGTCTavernInfoNotify, func() any { return new(proto.GCGTCTavernInfoNotify) })
	c.regMsg(GCGTavernNpcInfoNotify, func() any { return new(proto.GCGTavernNpcInfoNotify) })
	c.regMsg(GCGWeekChallengeInfoNotify, func() any { return new(proto.GCGWeekChallengeInfoNotify) })
	c.regMsg(GCGWorldChallengeUnlockNotify, func() any { return new(proto.GCGWorldChallengeUnlockNotify) })
	c.regMsg(GMShowNavMeshReq, func() any { return new(proto.GMShowNavMeshReq) })
	c.regMsg(GMShowNavMeshRsp, func() any { return new(proto.GMShowNavMeshRsp) })
	c.regMsg(GMShowObstacleReq, func() any { return new(proto.GMShowObstacleReq) })
	c.regMsg(GMShowObstacleRsp, func() any { return new(proto.GMShowObstacleRsp) })
	c.regMsg(GachaActivityCreateRobotReq, func() any { return new(proto.GachaActivityCreateRobotReq) })
	c.regMsg(GachaActivityCreateRobotRsp, func() any { return new(proto.GachaActivityCreateRobotRsp) })
	c.regMsg(GachaActivityNextStageReq, func() any { return new(proto.GachaActivityNextStageReq) })
	c.regMsg(GachaActivityNextStageRsp, func() any { return new(proto.GachaActivityNextStageRsp) })
	c.regMsg(GachaActivityPercentNotify, func() any { return new(proto.GachaActivityPercentNotify) })
	c.regMsg(GachaActivityResetReq, func() any { return new(proto.GachaActivityResetReq) })
	c.regMsg(GachaActivityResetRsp, func() any { return new(proto.GachaActivityResetRsp) })
	c.regMsg(GachaActivityTakeRewardReq, func() any { return new(proto.GachaActivityTakeRewardReq) })
	c.regMsg(GachaActivityTakeRewardRsp, func() any { return new(proto.GachaActivityTakeRewardRsp) })
	c.regMsg(GachaActivityUpdateElemNotify, func() any { return new(proto.GachaActivityUpdateElemNotify) })
	c.regMsg(GachaOpenWishNotify, func() any { return new(proto.GachaOpenWishNotify) })
	c.regMsg(GachaSimpleInfoNotify, func() any { return new(proto.GachaSimpleInfoNotify) })
	c.regMsg(GachaWishReq, func() any { return new(proto.GachaWishReq) })
	c.regMsg(GachaWishRsp, func() any { return new(proto.GachaWishRsp) })
	c.regMsg(GadgetAutoPickDropInfoNotify, func() any { return new(proto.GadgetAutoPickDropInfoNotify) })
	c.regMsg(GadgetChainLevelChangeNotify, func() any { return new(proto.GadgetChainLevelChangeNotify) })
	c.regMsg(GadgetChainLevelUpdateNotify, func() any { return new(proto.GadgetChainLevelUpdateNotify) })
	c.regMsg(GadgetChangeLevelTagReq, func() any { return new(proto.GadgetChangeLevelTagReq) })
	c.regMsg(GadgetChangeLevelTagRsp, func() any { return new(proto.GadgetChangeLevelTagRsp) })
	c.regMsg(GadgetCustomTreeInfoNotify, func() any { return new(proto.GadgetCustomTreeInfoNotify) })
	c.regMsg(GadgetGeneralRewardInfoNotify, func() any { return new(proto.GadgetGeneralRewardInfoNotify) })
	c.regMsg(GadgetInteractReq, func() any { return new(proto.GadgetInteractReq) })
	c.regMsg(GadgetInteractRsp, func() any { return new(proto.GadgetInteractRsp) })
	c.regMsg(GadgetPlayDataNotify, func() any { return new(proto.GadgetPlayDataNotify) })
	c.regMsg(GadgetPlayStartNotify, func() any { return new(proto.GadgetPlayStartNotify) })
	c.regMsg(GadgetPlayStopNotify, func() any { return new(proto.GadgetPlayStopNotify) })
	c.regMsg(GadgetPlayUidOpNotify, func() any { return new(proto.GadgetPlayUidOpNotify) })
	c.regMsg(GadgetStateNotify, func() any { return new(proto.GadgetStateNotify) })
	c.regMsg(GadgetTalkChangeNotify, func() any { return new(proto.GadgetTalkChangeNotify) })
	c.regMsg(GalleryBalloonScoreNotify, func() any { return new(proto.GalleryBalloonScoreNotify) })
	c.regMsg(GalleryBalloonShootNotify, func() any { return new(proto.GalleryBalloonShootNotify) })
	c.regMsg(GalleryBounceConjuringHitNotify, func() any { return new(proto.GalleryBounceConjuringHitNotify) })
	c.regMsg(GalleryBrokenFloorFallNotify, func() any { return new(proto.GalleryBrokenFloorFallNotify) })
	c.regMsg(GalleryBulletHitNotify, func() any { return new(proto.GalleryBulletHitNotify) })
	c.regMsg(GalleryCrystalLinkBuffInfoNotify, func() any { return new(proto.GalleryCrystalLinkBuffInfoNotify) })
	c.regMsg(GalleryCrystalLinkKillMonsterNotify, func() any { return new(proto.GalleryCrystalLinkKillMonsterNotify) })
	c.regMsg(GalleryFallCatchNotify, func() any { return new(proto.GalleryFallCatchNotify) })
	c.regMsg(GalleryFallScoreNotify, func() any { return new(proto.GalleryFallScoreNotify) })
	c.regMsg(GalleryFlowerCatchNotify, func() any { return new(proto.GalleryFlowerCatchNotify) })
	c.regMsg(GalleryIslandPartyDownHillInfoNotify, func() any { return new(proto.GalleryIslandPartyDownHillInfoNotify) })
	c.regMsg(GalleryPreStartNotify, func() any { return new(proto.GalleryPreStartNotify) })
	c.regMsg(GalleryStartNotify, func() any { return new(proto.GalleryStartNotify) })
	c.regMsg(GalleryStopNotify, func() any { return new(proto.GalleryStopNotify) })
	c.regMsg(GallerySumoKillMonsterNotify, func() any { return new(proto.GallerySumoKillMonsterNotify) })
	c.regMsg(GalleryWillStartCountdownNotify, func() any { return new(proto.GalleryWillStartCountdownNotify) })
	c.regMsg(GearActivityFinishPlayGearReq, func() any { return new(proto.GearActivityFinishPlayGearReq) })
	c.regMsg(GearActivityFinishPlayGearRsp, func() any { return new(proto.GearActivityFinishPlayGearRsp) })
	c.regMsg(GearActivityFinishPlayPictureReq, func() any { return new(proto.GearActivityFinishPlayPictureReq) })
	c.regMsg(GearActivityFinishPlayPictureRsp, func() any { return new(proto.GearActivityFinishPlayPictureRsp) })
	c.regMsg(GearActivityStartPlayGearReq, func() any { return new(proto.GearActivityStartPlayGearReq) })
	c.regMsg(GearActivityStartPlayGearRsp, func() any { return new(proto.GearActivityStartPlayGearRsp) })
	c.regMsg(GearActivityStartPlayPictureReq, func() any { return new(proto.GearActivityStartPlayPictureReq) })
	c.regMsg(GearActivityStartPlayPictureRsp, func() any { return new(proto.GearActivityStartPlayPictureRsp) })
	c.regMsg(GetActivityInfoReq, func() any { return new(proto.GetActivityInfoReq) })
	c.regMsg(GetActivityInfoRsp, func() any { return new(proto.GetActivityInfoRsp) })
	c.regMsg(GetActivityScheduleReq, func() any { return new(proto.GetActivityScheduleReq) })
	c.regMsg(GetActivityScheduleRsp, func() any { return new(proto.GetActivityScheduleRsp) })
	c.regMsg(GetActivityShopSheetInfoReq, func() any { return new(proto.GetActivityShopSheetInfoReq) })
	c.regMsg(GetActivityShopSheetInfoRsp, func() any { return new(proto.GetActivityShopSheetInfoRsp) })
	c.regMsg(GetAllActivatedBargainDataReq, func() any { return new(proto.GetAllActivatedBargainDataReq) })
	c.regMsg(GetAllActivatedBargainDataRsp, func() any { return new(proto.GetAllActivatedBargainDataRsp) })
	c.regMsg(GetAllH5ActivityInfoReq, func() any { return new(proto.GetAllH5ActivityInfoReq) })
	c.regMsg(GetAllH5ActivityInfoRsp, func() any { return new(proto.GetAllH5ActivityInfoRsp) })
	c.regMsg(GetAllMailNotify, func() any { return new(proto.GetAllMailNotify) })
	c.regMsg(GetAllMailReq, func() any { return new(proto.GetAllMailReq) })
	c.regMsg(GetAllMailResultNotify, func() any { return new(proto.GetAllMailResultNotify) })
	c.regMsg(GetAllMailRsp, func() any { return new(proto.GetAllMailRsp) })
	c.regMsg(GetAllSceneGalleryInfoReq, func() any { return new(proto.GetAllSceneGalleryInfoReq) })
	c.regMsg(GetAllSceneGalleryInfoRsp, func() any { return new(proto.GetAllSceneGalleryInfoRsp) })
	c.regMsg(GetAllUnlockNameCardReq, func() any { return new(proto.GetAllUnlockNameCardReq) })
	c.regMsg(GetAllUnlockNameCardRsp, func() any { return new(proto.GetAllUnlockNameCardRsp) })
	c.regMsg(GetAreaExplorePointReq, func() any { return new(proto.GetAreaExplorePointReq) })
	c.regMsg(GetAreaExplorePointRsp, func() any { return new(proto.GetAreaExplorePointRsp) })
	c.regMsg(GetAuthSalesmanInfoReq, func() any { return new(proto.GetAuthSalesmanInfoReq) })
	c.regMsg(GetAuthSalesmanInfoRsp, func() any { return new(proto.GetAuthSalesmanInfoRsp) })
	c.regMsg(GetAuthkeyReq, func() any { return new(proto.GetAuthkeyReq) })
	c.regMsg(GetAuthkeyRsp, func() any { return new(proto.GetAuthkeyRsp) })
	c.regMsg(GetBargainDataReq, func() any { return new(proto.GetBargainDataReq) })
	c.regMsg(GetBargainDataRsp, func() any { return new(proto.GetBargainDataRsp) })
	c.regMsg(GetBattlePassProductReq, func() any { return new(proto.GetBattlePassProductReq) })
	c.regMsg(GetBattlePassProductRsp, func() any { return new(proto.GetBattlePassProductRsp) })
	c.regMsg(GetBlossomBriefInfoListReq, func() any { return new(proto.GetBlossomBriefInfoListReq) })
	c.regMsg(GetBlossomBriefInfoListRsp, func() any { return new(proto.GetBlossomBriefInfoListRsp) })
	c.regMsg(GetBonusActivityRewardReq, func() any { return new(proto.GetBonusActivityRewardReq) })
	c.regMsg(GetBonusActivityRewardRsp, func() any { return new(proto.GetBonusActivityRewardRsp) })
	c.regMsg(GetChatEmojiCollectionReq, func() any { return new(proto.GetChatEmojiCollectionReq) })
	c.regMsg(GetChatEmojiCollectionRsp, func() any { return new(proto.GetChatEmojiCollectionRsp) })
	c.regMsg(GetCityHuntingOfferReq, func() any { return new(proto.GetCityHuntingOfferReq) })
	c.regMsg(GetCityHuntingOfferRsp, func() any { return new(proto.GetCityHuntingOfferRsp) })
	c.regMsg(GetCityReputationInfoReq, func() any { return new(proto.GetCityReputationInfoReq) })
	c.regMsg(GetCityReputationInfoRsp, func() any { return new(proto.GetCityReputationInfoRsp) })
	c.regMsg(GetCityReputationMapInfoReq, func() any { return new(proto.GetCityReputationMapInfoReq) })
	c.regMsg(GetCityReputationMapInfoRsp, func() any { return new(proto.GetCityReputationMapInfoRsp) })
	c.regMsg(GetCompoundDataReq, func() any { return new(proto.GetCompoundDataReq) })
	c.regMsg(GetCompoundDataRsp, func() any { return new(proto.GetCompoundDataRsp) })
	c.regMsg(GetCustomDungeonReq, func() any { return new(proto.GetCustomDungeonReq) })
	c.regMsg(GetCustomDungeonRsp, func() any { return new(proto.GetCustomDungeonRsp) })
	c.regMsg(GetDailyDungeonEntryInfoReq, func() any { return new(proto.GetDailyDungeonEntryInfoReq) })
	c.regMsg(GetDailyDungeonEntryInfoRsp, func() any { return new(proto.GetDailyDungeonEntryInfoRsp) })
	c.regMsg(GetDungeonEntryExploreConditionReq, func() any { return new(proto.GetDungeonEntryExploreConditionReq) })
	c.regMsg(GetDungeonEntryExploreConditionRsp, func() any { return new(proto.GetDungeonEntryExploreConditionRsp) })
	c.regMsg(GetExpeditionAssistInfoListReq, func() any { return new(proto.GetExpeditionAssistInfoListReq) })
	c.regMsg(GetExpeditionAssistInfoListRsp, func() any { return new(proto.GetExpeditionAssistInfoListRsp) })
	c.regMsg(GetFriendShowAvatarInfoReq, func() any { return new(proto.GetFriendShowAvatarInfoReq) })
	c.regMsg(GetFriendShowAvatarInfoRsp, func() any { return new(proto.GetFriendShowAvatarInfoRsp) })
	c.regMsg(GetFriendShowNameCardInfoReq, func() any { return new(proto.GetFriendShowNameCardInfoReq) })
	c.regMsg(GetFriendShowNameCardInfoRsp, func() any { return new(proto.GetFriendShowNameCardInfoRsp) })
	c.regMsg(GetFurnitureCurModuleArrangeCountReq, func() any { return new(proto.GetFurnitureCurModuleArrangeCountReq) })
	c.regMsg(GetGachaInfoReq, func() any { return new(proto.GetGachaInfoReq) })
	c.regMsg(GetGachaInfoRsp, func() any { return new(proto.GetGachaInfoRsp) })
	c.regMsg(GetGameplayRecommendationReq, func() any { return new(proto.GetGameplayRecommendationReq) })
	c.regMsg(GetGameplayRecommendationRsp, func() any { return new(proto.GetGameplayRecommendationRsp) })
	c.regMsg(GetHomeExchangeWoodInfoReq, func() any { return new(proto.GetHomeExchangeWoodInfoReq) })
	c.regMsg(GetHomeExchangeWoodInfoRsp, func() any { return new(proto.GetHomeExchangeWoodInfoRsp) })
	c.regMsg(GetHomeLevelUpRewardReq, func() any { return new(proto.GetHomeLevelUpRewardReq) })
	c.regMsg(GetHomeLevelUpRewardRsp, func() any { return new(proto.GetHomeLevelUpRewardRsp) })
	c.regMsg(GetHuntingOfferRewardReq, func() any { return new(proto.GetHuntingOfferRewardReq) })
	c.regMsg(GetHuntingOfferRewardRsp, func() any { return new(proto.GetHuntingOfferRewardRsp) })
	c.regMsg(GetInvestigationMonsterReq, func() any { return new(proto.GetInvestigationMonsterReq) })
	c.regMsg(GetInvestigationMonsterRsp, func() any { return new(proto.GetInvestigationMonsterRsp) })
	c.regMsg(GetMailItemReq, func() any { return new(proto.GetMailItemReq) })
	c.regMsg(GetMailItemRsp, func() any { return new(proto.GetMailItemRsp) })
	c.regMsg(GetMapAreaReq, func() any { return new(proto.GetMapAreaReq) })
	c.regMsg(GetMapAreaRsp, func() any { return new(proto.GetMapAreaRsp) })
	c.regMsg(GetMapMarkTipsReq, func() any { return new(proto.GetMapMarkTipsReq) })
	c.regMsg(GetMapMarkTipsRsp, func() any { return new(proto.GetMapMarkTipsRsp) })
	c.regMsg(GetMechanicusInfoReq, func() any { return new(proto.GetMechanicusInfoReq) })
	c.regMsg(GetMechanicusInfoRsp, func() any { return new(proto.GetMechanicusInfoRsp) })
	c.regMsg(GetNextResourceInfoReq, func() any { return new(proto.GetNextResourceInfoReq) })
	c.regMsg(GetNextResourceInfoRsp, func() any { return new(proto.GetNextResourceInfoRsp) })
	c.regMsg(GetOnlinePlayerInfoReq, func() any { return new(proto.GetOnlinePlayerInfoReq) })
	c.regMsg(GetOnlinePlayerInfoRsp, func() any { return new(proto.GetOnlinePlayerInfoRsp) })
	c.regMsg(GetOnlinePlayerListReq, func() any { return new(proto.GetOnlinePlayerListReq) })
	c.regMsg(GetOnlinePlayerListRsp, func() any { return new(proto.GetOnlinePlayerListRsp) })
	c.regMsg(GetOpActivityInfoReq, func() any { return new(proto.GetOpActivityInfoReq) })
	c.regMsg(GetOpActivityInfoRsp, func() any { return new(proto.GetOpActivityInfoRsp) })
	c.regMsg(GetParentQuestVideoKeyReq, func() any { return new(proto.GetParentQuestVideoKeyReq) })
	c.regMsg(GetParentQuestVideoKeyRsp, func() any { return new(proto.GetParentQuestVideoKeyRsp) })
	c.regMsg(GetPlayerAskFriendListReq, func() any { return new(proto.GetPlayerAskFriendListReq) })
	c.regMsg(GetPlayerAskFriendListRsp, func() any { return new(proto.GetPlayerAskFriendListRsp) })
	c.regMsg(GetPlayerBlacklistReq, func() any { return new(proto.GetPlayerBlacklistReq) })
	c.regMsg(GetPlayerBlacklistRsp, func() any { return new(proto.GetPlayerBlacklistRsp) })
	c.regMsg(GetPlayerFriendListReq, func() any { return new(proto.GetPlayerFriendListReq) })
	c.regMsg(GetPlayerFriendListRsp, func() any { return new(proto.GetPlayerFriendListRsp) })
	c.regMsg(GetPlayerHomeCompInfoReq, func() any { return new(proto.GetPlayerHomeCompInfoReq) })
	c.regMsg(GetPlayerMpModeAvailabilityReq, func() any { return new(proto.GetPlayerMpModeAvailabilityReq) })
	c.regMsg(GetPlayerMpModeAvailabilityRsp, func() any { return new(proto.GetPlayerMpModeAvailabilityRsp) })
	c.regMsg(GetPlayerSocialDetailReq, func() any { return new(proto.GetPlayerSocialDetailReq) })
	c.regMsg(GetPlayerSocialDetailRsp, func() any { return new(proto.GetPlayerSocialDetailRsp) })
	c.regMsg(GetPlayerTokenReq, func() any { return new(proto.GetPlayerTokenReq) })
	c.regMsg(GetPlayerTokenRsp, func() any { return new(proto.GetPlayerTokenRsp) })
	c.regMsg(GetPushTipsRewardReq, func() any { return new(proto.GetPushTipsRewardReq) })
	c.regMsg(GetPushTipsRewardRsp, func() any { return new(proto.GetPushTipsRewardRsp) })
	c.regMsg(GetQuestLackingResourceReq, func() any { return new(proto.GetQuestLackingResourceReq) })
	c.regMsg(GetQuestLackingResourceRsp, func() any { return new(proto.GetQuestLackingResourceRsp) })
	c.regMsg(GetQuestTalkHistoryReq, func() any { return new(proto.GetQuestTalkHistoryReq) })
	c.regMsg(GetQuestTalkHistoryRsp, func() any { return new(proto.GetQuestTalkHistoryRsp) })
	c.regMsg(GetRecentMpPlayerListReq, func() any { return new(proto.GetRecentMpPlayerListReq) })
	c.regMsg(GetRecentMpPlayerListRsp, func() any { return new(proto.GetRecentMpPlayerListRsp) })
	c.regMsg(GetRecommendCustomDungeonReq, func() any { return new(proto.GetRecommendCustomDungeonReq) })
	c.regMsg(GetRecommendCustomDungeonRsp, func() any { return new(proto.GetRecommendCustomDungeonRsp) })
	c.regMsg(GetRegionSearchReq, func() any { return new(proto.GetRegionSearchReq) })
	c.regMsg(GetReunionMissionInfoReq, func() any { return new(proto.GetReunionMissionInfoReq) })
	c.regMsg(GetReunionMissionInfoRsp, func() any { return new(proto.GetReunionMissionInfoRsp) })
	c.regMsg(GetReunionPrivilegeInfoReq, func() any { return new(proto.GetReunionPrivilegeInfoReq) })
	c.regMsg(GetReunionPrivilegeInfoRsp, func() any { return new(proto.GetReunionPrivilegeInfoRsp) })
	c.regMsg(GetReunionSignInInfoReq, func() any { return new(proto.GetReunionSignInInfoReq) })
	c.regMsg(GetReunionSignInInfoRsp, func() any { return new(proto.GetReunionSignInInfoRsp) })
	c.regMsg(GetRogueDairyRepairInfoReq, func() any { return new(proto.GetRogueDairyRepairInfoReq) })
	c.regMsg(GetRogueDairyRepairInfoRsp, func() any { return new(proto.GetRogueDairyRepairInfoRsp) })
	c.regMsg(GetSceneAreaReq, func() any { return new(proto.GetSceneAreaReq) })
	c.regMsg(GetSceneAreaRsp, func() any { return new(proto.GetSceneAreaRsp) })
	c.regMsg(GetSceneNpcPositionReq, func() any { return new(proto.GetSceneNpcPositionReq) })
	c.regMsg(GetSceneNpcPositionRsp, func() any { return new(proto.GetSceneNpcPositionRsp) })
	c.regMsg(GetScenePerformanceReq, func() any { return new(proto.GetScenePerformanceReq) })
	c.regMsg(GetScenePerformanceRsp, func() any { return new(proto.GetScenePerformanceRsp) })
	c.regMsg(GetScenePointReq, func() any { return new(proto.GetScenePointReq) })
	c.regMsg(GetScenePointRsp, func() any { return new(proto.GetScenePointRsp) })
	c.regMsg(GetShopReq, func() any { return new(proto.GetShopReq) })
	c.regMsg(GetShopRsp, func() any { return new(proto.GetShopRsp) })
	c.regMsg(GetShopmallDataReq, func() any { return new(proto.GetShopmallDataReq) })
	c.regMsg(GetShopmallDataRsp, func() any { return new(proto.GetShopmallDataRsp) })
	c.regMsg(GetSignInRewardReq, func() any { return new(proto.GetSignInRewardReq) })
	c.regMsg(GetSignInRewardRsp, func() any { return new(proto.GetSignInRewardRsp) })
	c.regMsg(GetStoreCustomDungeonReq, func() any { return new(proto.GetStoreCustomDungeonReq) })
	c.regMsg(GetStoreCustomDungeonRsp, func() any { return new(proto.GetStoreCustomDungeonRsp) })
	c.regMsg(GetUgcBriefInfoReq, func() any { return new(proto.GetUgcBriefInfoReq) })
	c.regMsg(GetUgcBriefInfoRsp, func() any { return new(proto.GetUgcBriefInfoRsp) })
	c.regMsg(GetUgcReq, func() any { return new(proto.GetUgcReq) })
	c.regMsg(GetUgcRsp, func() any { return new(proto.GetUgcRsp) })
	c.regMsg(GetWidgetSlotReq, func() any { return new(proto.GetWidgetSlotReq) })
	c.regMsg(GetWidgetSlotRsp, func() any { return new(proto.GetWidgetSlotRsp) })
	c.regMsg(GetWorldMpInfoReq, func() any { return new(proto.GetWorldMpInfoReq) })
	c.regMsg(GetWorldMpInfoRsp, func() any { return new(proto.GetWorldMpInfoRsp) })
	c.regMsg(GiveUpRoguelikeDungeonCardReq, func() any { return new(proto.GiveUpRoguelikeDungeonCardReq) })
	c.regMsg(GiveUpRoguelikeDungeonCardRsp, func() any { return new(proto.GiveUpRoguelikeDungeonCardRsp) })
	c.regMsg(GivingRecordChangeNotify, func() any { return new(proto.GivingRecordChangeNotify) })
	c.regMsg(GivingRecordNotify, func() any { return new(proto.GivingRecordNotify) })
	c.regMsg(GlobalBuildingInfoNotify, func() any { return new(proto.GlobalBuildingInfoNotify) })
	c.regMsg(GmTalkNotify, func() any { return new(proto.GmTalkNotify) })
	c.regMsg(GmTalkReq, func() any { return new(proto.GmTalkReq) })
	c.regMsg(GmTalkRsp, func() any { return new(proto.GmTalkRsp) })
	c.regMsg(GrantRewardNotify, func() any { return new(proto.GrantRewardNotify) })
	c.regMsg(GravenInnocenceEditCarveCombinationReq, func() any { return new(proto.GravenInnocenceEditCarveCombinationReq) })
	c.regMsg(GravenInnocenceEditCarveCombinationRsp, func() any { return new(proto.GravenInnocenceEditCarveCombinationRsp) })
	c.regMsg(GravenInnocencePhotoFinishReq, func() any { return new(proto.GravenInnocencePhotoFinishReq) })
	c.regMsg(GravenInnocencePhotoFinishRsp, func() any { return new(proto.GravenInnocencePhotoFinishRsp) })
	c.regMsg(GravenInnocencePhotoReminderNotify, func() any { return new(proto.GravenInnocencePhotoReminderNotify) })
	c.regMsg(GravenInnocenceRaceRestartReq, func() any { return new(proto.GravenInnocenceRaceRestartReq) })
	c.regMsg(GravenInnocenceRaceRestartRsp, func() any { return new(proto.GravenInnocenceRaceRestartRsp) })
	c.regMsg(GravenInnocenceRaceSettleNotify, func() any { return new(proto.GravenInnocenceRaceSettleNotify) })
	c.regMsg(GroupLinkAllNotify, func() any { return new(proto.GroupLinkAllNotify) })
	c.regMsg(GroupLinkChangeNotify, func() any { return new(proto.GroupLinkChangeNotify) })
	c.regMsg(GroupLinkDeleteNotify, func() any { return new(proto.GroupLinkDeleteNotify) })
	c.regMsg(GroupLinkMarkUpdateNotify, func() any { return new(proto.GroupLinkMarkUpdateNotify) })
	c.regMsg(GroupSuiteNotify, func() any { return new(proto.GroupSuiteNotify) })
	c.regMsg(GroupUnloadNotify, func() any { return new(proto.GroupUnloadNotify) })
	c.regMsg(GuestBeginEnterSceneNotify, func() any { return new(proto.GuestBeginEnterSceneNotify) })
	c.regMsg(GuestPostEnterSceneNotify, func() any { return new(proto.GuestPostEnterSceneNotify) })
	c.regMsg(H5ActivityIdsNotify, func() any { return new(proto.H5ActivityIdsNotify) })
	c.regMsg(HideAndSeekPlayerReadyNotify, func() any { return new(proto.HideAndSeekPlayerReadyNotify) })
	c.regMsg(HideAndSeekPlayerSetAvatarNotify, func() any { return new(proto.HideAndSeekPlayerSetAvatarNotify) })
	c.regMsg(HideAndSeekSelectAvatarReq, func() any { return new(proto.HideAndSeekSelectAvatarReq) })
	c.regMsg(HideAndSeekSelectAvatarRsp, func() any { return new(proto.HideAndSeekSelectAvatarRsp) })
	c.regMsg(HideAndSeekSelectSkillReq, func() any { return new(proto.HideAndSeekSelectSkillReq) })
	c.regMsg(HideAndSeekSelectSkillRsp, func() any { return new(proto.HideAndSeekSelectSkillRsp) })
	c.regMsg(HideAndSeekSetReadyReq, func() any { return new(proto.HideAndSeekSetReadyReq) })
	c.regMsg(HideAndSeekSetReadyRsp, func() any { return new(proto.HideAndSeekSetReadyRsp) })
	c.regMsg(HideAndSeekSettleNotify, func() any { return new(proto.HideAndSeekSettleNotify) })
	c.regMsg(HitClientTrivialNotify, func() any { return new(proto.HitClientTrivialNotify) })
	c.regMsg(HitTreeNotify, func() any { return new(proto.HitTreeNotify) })
	c.regMsg(HomeAllUnlockedBgmIdListNotify, func() any { return new(proto.HomeAllUnlockedBgmIdListNotify) })
	c.regMsg(HomeAvatarAllFinishRewardNotify, func() any { return new(proto.HomeAvatarAllFinishRewardNotify) })
	c.regMsg(HomeAvatarCostumeChangeNotify, func() any { return new(proto.HomeAvatarCostumeChangeNotify) })
	c.regMsg(HomeAvatarRewardEventGetReq, func() any { return new(proto.HomeAvatarRewardEventGetReq) })
	c.regMsg(HomeAvatarRewardEventGetRsp, func() any { return new(proto.HomeAvatarRewardEventGetRsp) })
	c.regMsg(HomeAvatarRewardEventNotify, func() any { return new(proto.HomeAvatarRewardEventNotify) })
	c.regMsg(HomeAvatarSummonAllEventNotify, func() any { return new(proto.HomeAvatarSummonAllEventNotify) })
	c.regMsg(HomeAvatarSummonEventReq, func() any { return new(proto.HomeAvatarSummonEventReq) })
	c.regMsg(HomeAvatarSummonEventRsp, func() any { return new(proto.HomeAvatarSummonEventRsp) })
	c.regMsg(HomeAvatarSummonFinishReq, func() any { return new(proto.HomeAvatarSummonFinishReq) })
	c.regMsg(HomeAvatarSummonFinishRsp, func() any { return new(proto.HomeAvatarSummonFinishRsp) })
	c.regMsg(HomeAvatarTalkFinishInfoNotify, func() any { return new(proto.HomeAvatarTalkFinishInfoNotify) })
	c.regMsg(HomeAvatarTalkReq, func() any { return new(proto.HomeAvatarTalkReq) })
	c.regMsg(HomeAvatarTalkRsp, func() any { return new(proto.HomeAvatarTalkRsp) })
	c.regMsg(HomeAvtarAllFinishRewardNotify, func() any { return new(proto.HomeAvtarAllFinishRewardNotify) })
	c.regMsg(HomeBalloonGalleryScoreNotify, func() any { return new(proto.HomeBalloonGalleryScoreNotify) })
	c.regMsg(HomeBalloonGallerySettleNotify, func() any { return new(proto.HomeBalloonGallerySettleNotify) })
	c.regMsg(HomeBasicInfoNotify, func() any { return new(proto.HomeBasicInfoNotify) })
	c.regMsg(HomeBlockNotify, func() any { return new(proto.HomeBlockNotify) })
	c.regMsg(HomeBlueprintInfoNotify, func() any { return new(proto.HomeBlueprintInfoNotify) })
	c.regMsg(HomeChangeBgmNotify, func() any { return new(proto.HomeChangeBgmNotify) })
	c.regMsg(HomeChangeBgmReq, func() any { return new(proto.HomeChangeBgmReq) })
	c.regMsg(HomeChangeBgmRsp, func() any { return new(proto.HomeChangeBgmRsp) })
	c.regMsg(HomeChangeEditModeReq, func() any { return new(proto.HomeChangeEditModeReq) })
	c.regMsg(HomeChangeEditModeRsp, func() any { return new(proto.HomeChangeEditModeRsp) })
	c.regMsg(HomeChangeModuleReq, func() any { return new(proto.HomeChangeModuleReq) })
	c.regMsg(HomeChangeModuleRsp, func() any { return new(proto.HomeChangeModuleRsp) })
	c.regMsg(HomeChooseModuleReq, func() any { return new(proto.HomeChooseModuleReq) })
	c.regMsg(HomeChooseModuleRsp, func() any { return new(proto.HomeChooseModuleRsp) })
	c.regMsg(HomeClearGroupRecordReq, func() any { return new(proto.HomeClearGroupRecordReq) })
	c.regMsg(HomeClearGroupRecordRsp, func() any { return new(proto.HomeClearGroupRecordRsp) })
	c.regMsg(HomeComfortInfoNotify, func() any { return new(proto.HomeComfortInfoNotify) })
	c.regMsg(HomeCreateBlueprintReq, func() any { return new(proto.HomeCreateBlueprintReq) })
	c.regMsg(HomeCreateBlueprintRsp, func() any { return new(proto.HomeCreateBlueprintRsp) })
	c.regMsg(HomeCustomFurnitureInfoNotify, func() any { return new(proto.HomeCustomFurnitureInfoNotify) })
	c.regMsg(HomeDeleteBlueprintReq, func() any { return new(proto.HomeDeleteBlueprintReq) })
	c.regMsg(HomeDeleteBlueprintRsp, func() any { return new(proto.HomeDeleteBlueprintRsp) })
	c.regMsg(HomeEditCustomFurnitureReq, func() any { return new(proto.HomeEditCustomFurnitureReq) })
	c.regMsg(HomeEditCustomFurnitureRsp, func() any { return new(proto.HomeEditCustomFurnitureRsp) })
	c.regMsg(HomeEnterEditModeFinishReq, func() any { return new(proto.HomeEnterEditModeFinishReq) })
	c.regMsg(HomeEnterEditModeFinishRsp, func() any { return new(proto.HomeEnterEditModeFinishRsp) })
	c.regMsg(HomeExchangeWoodReq, func() any { return new(proto.HomeExchangeWoodReq) })
	c.regMsg(HomeExchangeWoodRsp, func() any { return new(proto.HomeExchangeWoodRsp) })
	c.regMsg(HomeFishFarmingInfoNotify, func() any { return new(proto.HomeFishFarmingInfoNotify) })
	c.regMsg(HomeGalleryInPlayingNotify, func() any { return new(proto.HomeGalleryInPlayingNotify) })
	c.regMsg(HomeGetArrangementInfoReq, func() any { return new(proto.HomeGetArrangementInfoReq) })
	c.regMsg(HomeGetArrangementInfoRsp, func() any { return new(proto.HomeGetArrangementInfoRsp) })
	c.regMsg(HomeGetBasicInfoReq, func() any { return new(proto.HomeGetBasicInfoReq) })
	c.regMsg(HomeGetBlueprintSlotInfoReq, func() any { return new(proto.HomeGetBlueprintSlotInfoReq) })
	c.regMsg(HomeGetBlueprintSlotInfoRsp, func() any { return new(proto.HomeGetBlueprintSlotInfoRsp) })
	c.regMsg(HomeGetFishFarmingInfoReq, func() any { return new(proto.HomeGetFishFarmingInfoReq) })
	c.regMsg(HomeGetFishFarmingInfoRsp, func() any { return new(proto.HomeGetFishFarmingInfoRsp) })
	c.regMsg(HomeGetGroupRecordReq, func() any { return new(proto.HomeGetGroupRecordReq) })
	c.regMsg(HomeGetGroupRecordRsp, func() any { return new(proto.HomeGetGroupRecordRsp) })
	c.regMsg(HomeGetOnlineStatusReq, func() any { return new(proto.HomeGetOnlineStatusReq) })
	c.regMsg(HomeGetOnlineStatusRsp, func() any { return new(proto.HomeGetOnlineStatusRsp) })
	c.regMsg(HomeKickPlayerReq, func() any { return new(proto.HomeKickPlayerReq) })
	c.regMsg(HomeKickPlayerRsp, func() any { return new(proto.HomeKickPlayerRsp) })
	c.regMsg(HomeLimitedShopBuyGoodsReq, func() any { return new(proto.HomeLimitedShopBuyGoodsReq) })
	c.regMsg(HomeLimitedShopBuyGoodsRsp, func() any { return new(proto.HomeLimitedShopBuyGoodsRsp) })
	c.regMsg(HomeLimitedShopGoodsListReq, func() any { return new(proto.HomeLimitedShopGoodsListReq) })
	c.regMsg(HomeLimitedShopGoodsListRsp, func() any { return new(proto.HomeLimitedShopGoodsListRsp) })
	c.regMsg(HomeLimitedShopInfoChangeNotify, func() any { return new(proto.HomeLimitedShopInfoChangeNotify) })
	c.regMsg(HomeLimitedShopInfoNotify, func() any { return new(proto.HomeLimitedShopInfoNotify) })
	c.regMsg(HomeLimitedShopInfoReq, func() any { return new(proto.HomeLimitedShopInfoReq) })
	c.regMsg(HomeLimitedShopInfoRsp, func() any { return new(proto.HomeLimitedShopInfoRsp) })
	c.regMsg(HomeMarkPointNotify, func() any { return new(proto.HomeMarkPointNotify) })
	c.regMsg(HomeModuleSeenReq, func() any { return new(proto.HomeModuleSeenReq) })
	c.regMsg(HomeModuleSeenRsp, func() any { return new(proto.HomeModuleSeenRsp) })
	c.regMsg(HomeModuleUnlockNotify, func() any { return new(proto.HomeModuleUnlockNotify) })
	c.regMsg(HomeNewUnlockedBgmIdListNotify, func() any { return new(proto.HomeNewUnlockedBgmIdListNotify) })
	c.regMsg(HomePictureFrameInfoNotify, func() any { return new(proto.HomePictureFrameInfoNotify) })
	c.regMsg(HomePlantFieldNotify, func() any { return new(proto.HomePlantFieldNotify) })
	c.regMsg(HomePlantInfoNotify, func() any { return new(proto.HomePlantInfoNotify) })
	c.regMsg(HomePlantInfoReq, func() any { return new(proto.HomePlantInfoReq) })
	c.regMsg(HomePlantInfoRsp, func() any { return new(proto.HomePlantInfoRsp) })
	c.regMsg(HomePlantSeedReq, func() any { return new(proto.HomePlantSeedReq) })
	c.regMsg(HomePlantSeedRsp, func() any { return new(proto.HomePlantSeedRsp) })
	c.regMsg(HomePlantWeedReq, func() any { return new(proto.HomePlantWeedReq) })
	c.regMsg(HomePlantWeedRsp, func() any { return new(proto.HomePlantWeedRsp) })
	c.regMsg(HomePreChangeEditModeNotify, func() any { return new(proto.HomePreChangeEditModeNotify) })
	c.regMsg(HomePreviewBlueprintReq, func() any { return new(proto.HomePreviewBlueprintReq) })
	c.regMsg(HomePreviewBlueprintRsp, func() any { return new(proto.HomePreviewBlueprintRsp) })
	c.regMsg(HomePriorCheckNotify, func() any { return new(proto.HomePriorCheckNotify) })
	c.regMsg(HomeRacingGallerySettleNotify, func() any { return new(proto.HomeRacingGallerySettleNotify) })
	c.regMsg(HomeResourceNotify, func() any { return new(proto.HomeResourceNotify) })
	c.regMsg(HomeResourceTakeFetterExpReq, func() any { return new(proto.HomeResourceTakeFetterExpReq) })
	c.regMsg(HomeResourceTakeFetterExpRsp, func() any { return new(proto.HomeResourceTakeFetterExpRsp) })
	c.regMsg(HomeResourceTakeHomeCoinReq, func() any { return new(proto.HomeResourceTakeHomeCoinReq) })
	c.regMsg(HomeResourceTakeHomeCoinRsp, func() any { return new(proto.HomeResourceTakeHomeCoinRsp) })
	c.regMsg(HomeSaveArrangementNoChangeReq, func() any { return new(proto.HomeSaveArrangementNoChangeReq) })
	c.regMsg(HomeSaveArrangementNoChangeRsp, func() any { return new(proto.HomeSaveArrangementNoChangeRsp) })
	c.regMsg(HomeSceneInitFinishReq, func() any { return new(proto.HomeSceneInitFinishReq) })
	c.regMsg(HomeSceneInitFinishRsp, func() any { return new(proto.HomeSceneInitFinishRsp) })
	c.regMsg(HomeSceneJumpReq, func() any { return new(proto.HomeSceneJumpReq) })
	c.regMsg(HomeSceneJumpRsp, func() any { return new(proto.HomeSceneJumpRsp) })
	c.regMsg(HomeScenePointFishFarmingInfoNotify, func() any { return new(proto.HomeScenePointFishFarmingInfoNotify) })
	c.regMsg(HomeSearchBlueprintReq, func() any { return new(proto.HomeSearchBlueprintReq) })
	c.regMsg(HomeSearchBlueprintRsp, func() any { return new(proto.HomeSearchBlueprintRsp) })
	c.regMsg(HomeSeekFurnitureGalleryScoreNotify, func() any { return new(proto.HomeSeekFurnitureGalleryScoreNotify) })
	c.regMsg(HomeSetBlueprintFriendOptionReq, func() any { return new(proto.HomeSetBlueprintFriendOptionReq) })
	c.regMsg(HomeSetBlueprintFriendOptionRsp, func() any { return new(proto.HomeSetBlueprintFriendOptionRsp) })
	c.regMsg(HomeSetBlueprintSlotOptionReq, func() any { return new(proto.HomeSetBlueprintSlotOptionReq) })
	c.regMsg(HomeSetBlueprintSlotOptionRsp, func() any { return new(proto.HomeSetBlueprintSlotOptionRsp) })
	c.regMsg(HomeTransferReq, func() any { return new(proto.HomeTransferReq) })
	c.regMsg(HomeTransferRsp, func() any { return new(proto.HomeTransferRsp) })
	c.regMsg(HomeUpdateArrangementInfoReq, func() any { return new(proto.HomeUpdateArrangementInfoReq) })
	c.regMsg(HomeUpdateArrangementInfoRsp, func() any { return new(proto.HomeUpdateArrangementInfoRsp) })
	c.regMsg(HomeUpdateFishFarmingInfoReq, func() any { return new(proto.HomeUpdateFishFarmingInfoReq) })
	c.regMsg(HomeUpdateFishFarmingInfoRsp, func() any { return new(proto.HomeUpdateFishFarmingInfoRsp) })
	c.regMsg(HomeUpdatePictureFrameInfoReq, func() any { return new(proto.HomeUpdatePictureFrameInfoReq) })
	c.regMsg(HomeUpdatePictureFrameInfoRsp, func() any { return new(proto.HomeUpdatePictureFrameInfoRsp) })
	c.regMsg(HomeUpdateScenePointFishFarmingInfoReq, func() any { return new(proto.HomeUpdateScenePointFishFarmingInfoReq) })
	c.regMsg(HomeUpdateScenePointFishFarmingInfoRsp, func() any { return new(proto.HomeUpdateScenePointFishFarmingInfoRsp) })
	c.regMsg(HostPlayerNotify, func() any { return new(proto.HostPlayerNotify) })
	c.regMsg(HuntingFailNotify, func() any { return new(proto.HuntingFailNotify) })
	c.regMsg(HuntingGiveUpReq, func() any { return new(proto.HuntingGiveUpReq) })
	c.regMsg(HuntingGiveUpRsp, func() any { return new(proto.HuntingGiveUpRsp) })
	c.regMsg(HuntingOngoingNotify, func() any { return new(proto.HuntingOngoingNotify) })
	c.regMsg(HuntingRevealClueNotify, func() any { return new(proto.HuntingRevealClueNotify) })
	c.regMsg(HuntingRevealFinalNotify, func() any { return new(proto.HuntingRevealFinalNotify) })
	c.regMsg(HuntingStartNotify, func() any { return new(proto.HuntingStartNotify) })
	c.regMsg(HuntingSuccessNotify, func() any { return new(proto.HuntingSuccessNotify) })
	c.regMsg(InBattleMechanicusBuildingPointsNotify, func() any { return new(proto.InBattleMechanicusBuildingPointsNotify) })
	c.regMsg(InBattleMechanicusCardResultNotify, func() any { return new(proto.InBattleMechanicusCardResultNotify) })
	c.regMsg(InBattleMechanicusConfirmCardNotify, func() any { return new(proto.InBattleMechanicusConfirmCardNotify) })
	c.regMsg(InBattleMechanicusConfirmCardReq, func() any { return new(proto.InBattleMechanicusConfirmCardReq) })
	c.regMsg(InBattleMechanicusConfirmCardRsp, func() any { return new(proto.InBattleMechanicusConfirmCardRsp) })
	c.regMsg(InBattleMechanicusLeftMonsterNotify, func() any { return new(proto.InBattleMechanicusLeftMonsterNotify) })
	c.regMsg(InBattleMechanicusPickCardNotify, func() any { return new(proto.InBattleMechanicusPickCardNotify) })
	c.regMsg(InBattleMechanicusPickCardReq, func() any { return new(proto.InBattleMechanicusPickCardReq) })
	c.regMsg(InBattleMechanicusPickCardRsp, func() any { return new(proto.InBattleMechanicusPickCardRsp) })
	c.regMsg(InBattleMechanicusSettleNotify, func() any { return new(proto.InBattleMechanicusSettleNotify) })
	c.regMsg(InstableSprayEnterDungeonReq, func() any { return new(proto.InstableSprayEnterDungeonReq) })
	c.regMsg(InstableSprayEnterDungeonRsp, func() any { return new(proto.InstableSprayEnterDungeonRsp) })
	c.regMsg(InstableSprayGalleryInfoNotify, func() any { return new(proto.InstableSprayGalleryInfoNotify) })
	c.regMsg(InstableSprayLevelFinishNotify, func() any { return new(proto.InstableSprayLevelFinishNotify) })
	c.regMsg(InstableSprayRestartDungeonReq, func() any { return new(proto.InstableSprayRestartDungeonReq) })
	c.regMsg(InstableSprayRestartDungeonRsp, func() any { return new(proto.InstableSprayRestartDungeonRsp) })
	c.regMsg(InstableSpraySwitchTeamReq, func() any { return new(proto.InstableSpraySwitchTeamReq) })
	c.regMsg(InstableSpraySwitchTeamRsp, func() any { return new(proto.InstableSpraySwitchTeamRsp) })
	c.regMsg(InteractDailyDungeonInfoNotify, func() any { return new(proto.InteractDailyDungeonInfoNotify) })
	c.regMsg(InterpretInferenceWordReq, func() any { return new(proto.InterpretInferenceWordReq) })
	c.regMsg(InterpretInferenceWordRsp, func() any { return new(proto.InterpretInferenceWordRsp) })
	c.regMsg(InterruptGalleryReq, func() any { return new(proto.InterruptGalleryReq) })
	c.regMsg(InterruptGalleryRsp, func() any { return new(proto.InterruptGalleryRsp) })
	c.regMsg(InvestigationMonsterUpdateNotify, func() any { return new(proto.InvestigationMonsterUpdateNotify) })
	c.regMsg(InvestigationQuestDailyNotify, func() any { return new(proto.InvestigationQuestDailyNotify) })
	c.regMsg(InvestigationReadQuestDailyNotify, func() any { return new(proto.InvestigationReadQuestDailyNotify) })
	c.regMsg(IrodoriChessEquipCardReq, func() any { return new(proto.IrodoriChessEquipCardReq) })
	c.regMsg(IrodoriChessEquipCardRsp, func() any { return new(proto.IrodoriChessEquipCardRsp) })
	c.regMsg(IrodoriChessLeftMonsterNotify, func() any { return new(proto.IrodoriChessLeftMonsterNotify) })
	c.regMsg(IrodoriChessPlayerInfoNotify, func() any { return new(proto.IrodoriChessPlayerInfoNotify) })
	c.regMsg(IrodoriChessUnequipCardReq, func() any { return new(proto.IrodoriChessUnequipCardReq) })
	c.regMsg(IrodoriChessUnequipCardRsp, func() any { return new(proto.IrodoriChessUnequipCardRsp) })
	c.regMsg(IrodoriEditFlowerCombinationReq, func() any { return new(proto.IrodoriEditFlowerCombinationReq) })
	c.regMsg(IrodoriEditFlowerCombinationRsp, func() any { return new(proto.IrodoriEditFlowerCombinationRsp) })
	c.regMsg(IrodoriFillPoetryReq, func() any { return new(proto.IrodoriFillPoetryReq) })
	c.regMsg(IrodoriFillPoetryRsp, func() any { return new(proto.IrodoriFillPoetryRsp) })
	c.regMsg(IrodoriMasterGalleryCgEndNotify, func() any { return new(proto.IrodoriMasterGalleryCgEndNotify) })
	c.regMsg(IrodoriMasterGallerySettleNotify, func() any { return new(proto.IrodoriMasterGallerySettleNotify) })
	c.regMsg(IrodoriMasterStartGalleryReq, func() any { return new(proto.IrodoriMasterStartGalleryReq) })
	c.regMsg(IrodoriMasterStartGalleryRsp, func() any { return new(proto.IrodoriMasterStartGalleryRsp) })
	c.regMsg(IrodoriScanEntityReq, func() any { return new(proto.IrodoriScanEntityReq) })
	c.regMsg(IrodoriScanEntityRsp, func() any { return new(proto.IrodoriScanEntityRsp) })
	c.regMsg(IslandPartyRaftInfoNotify, func() any { return new(proto.IslandPartyRaftInfoNotify) })
	c.regMsg(IslandPartySailInfoNotify, func() any { return new(proto.IslandPartySailInfoNotify) })
	c.regMsg(IslandPartySettleNotify, func() any { return new(proto.IslandPartySettleNotify) })
	c.regMsg(ItemAddHintNotify, func() any { return new(proto.ItemAddHintNotify) })
	c.regMsg(ItemCdGroupTimeNotify, func() any { return new(proto.ItemCdGroupTimeNotify) })
	c.regMsg(ItemGivingReq, func() any { return new(proto.ItemGivingReq) })
	c.regMsg(ItemGivingRsp, func() any { return new(proto.ItemGivingRsp) })
	c.regMsg(JoinHomeWorldFailNotify, func() any { return new(proto.JoinHomeWorldFailNotify) })
	c.regMsg(JoinPlayerFailNotify, func() any { return new(proto.JoinPlayerFailNotify) })
	c.regMsg(JoinPlayerSceneReq, func() any { return new(proto.JoinPlayerSceneReq) })
	c.regMsg(JoinPlayerSceneRsp, func() any { return new(proto.JoinPlayerSceneRsp) })
	c.regMsg(KeepAliveNotify, func() any { return new(proto.KeepAliveNotify) })
	c.regMsg(LanternRiteDoFireworksReformReq, func() any { return new(proto.LanternRiteDoFireworksReformReq) })
	c.regMsg(LanternRiteDoFireworksReformRsp, func() any { return new(proto.LanternRiteDoFireworksReformRsp) })
	c.regMsg(LanternRiteEndFireworksReformReq, func() any { return new(proto.LanternRiteEndFireworksReformReq) })
	c.regMsg(LanternRiteEndFireworksReformRsp, func() any { return new(proto.LanternRiteEndFireworksReformRsp) })
	c.regMsg(LanternRiteStartFireworksReformReq, func() any { return new(proto.LanternRiteStartFireworksReformReq) })
	c.regMsg(LanternRiteStartFireworksReformRsp, func() any { return new(proto.LanternRiteStartFireworksReformRsp) })
	c.regMsg(LanternRiteTakeSkinRewardReq, func() any { return new(proto.LanternRiteTakeSkinRewardReq) })
	c.regMsg(LanternRiteTakeSkinRewardRsp, func() any { return new(proto.LanternRiteTakeSkinRewardRsp) })
	c.regMsg(LastPacketPrintNotify, func() any { return new(proto.LastPacketPrintNotify) })
	c.regMsg(LaunchFireworksReq, func() any { return new(proto.LaunchFireworksReq) })
	c.regMsg(LaunchFireworksRsp, func() any { return new(proto.LaunchFireworksRsp) })
	c.regMsg(LeaveSceneReq, func() any { return new(proto.LeaveSceneReq) })
	c.regMsg(LeaveSceneRsp, func() any { return new(proto.LeaveSceneRsp) })
	c.regMsg(LeaveWorldNotify, func() any { return new(proto.LeaveWorldNotify) })
	c.regMsg(LevelTagDataNotify, func() any { return new(proto.LevelTagDataNotify) })
	c.regMsg(LevelupCityReq, func() any { return new(proto.LevelupCityReq) })
	c.regMsg(LevelupCityRsp, func() any { return new(proto.LevelupCityRsp) })
	c.regMsg(LifeStateChangeNotify, func() any { return new(proto.LifeStateChangeNotify) })
	c.regMsg(LikeCustomDungeonReq, func() any { return new(proto.LikeCustomDungeonReq) })
	c.regMsg(LikeCustomDungeonRsp, func() any { return new(proto.LikeCustomDungeonRsp) })
	c.regMsg(LiveEndNotify, func() any { return new(proto.LiveEndNotify) })
	c.regMsg(LiveStartNotify, func() any { return new(proto.LiveStartNotify) })
	c.regMsg(LoadActivityTerrainNotify, func() any { return new(proto.LoadActivityTerrainNotify) })
	c.regMsg(LuaEnvironmentEffectNotify, func() any { return new(proto.LuaEnvironmentEffectNotify) })
	c.regMsg(LuaSetOptionNotify, func() any { return new(proto.LuaSetOptionNotify) })
	c.regMsg(LuminanceStoneChallengeSettleNotify, func() any { return new(proto.LuminanceStoneChallengeSettleNotify) })
	c.regMsg(LunaRiteAreaFinishNotify, func() any { return new(proto.LunaRiteAreaFinishNotify) })
	c.regMsg(LunaRiteGroupBundleRegisterNotify, func() any { return new(proto.LunaRiteGroupBundleRegisterNotify) })
	c.regMsg(LunaRiteHintPointRemoveNotify, func() any { return new(proto.LunaRiteHintPointRemoveNotify) })
	c.regMsg(LunaRiteHintPointReq, func() any { return new(proto.LunaRiteHintPointReq) })
	c.regMsg(LunaRiteHintPointRsp, func() any { return new(proto.LunaRiteHintPointRsp) })
	c.regMsg(LunaRiteSacrificeReq, func() any { return new(proto.LunaRiteSacrificeReq) })
	c.regMsg(LunaRiteSacrificeRsp, func() any { return new(proto.LunaRiteSacrificeRsp) })
	c.regMsg(LunaRiteTakeSacrificeRewardReq, func() any { return new(proto.LunaRiteTakeSacrificeRewardReq) })
	c.regMsg(LunaRiteTakeSacrificeRewardRsp, func() any { return new(proto.LunaRiteTakeSacrificeRewardRsp) })
	c.regMsg(MailChangeNotify, func() any { return new(proto.MailChangeNotify) })
	c.regMsg(MainCoopFailNotify, func() any { return new(proto.MainCoopFailNotify) })
	c.regMsg(MainCoopUpdateNotify, func() any { return new(proto.MainCoopUpdateNotify) })
	c.regMsg(MapAreaChangeNotify, func() any { return new(proto.MapAreaChangeNotify) })
	c.regMsg(MarkEntityInMinMapNotify, func() any { return new(proto.MarkEntityInMinMapNotify) })
	c.regMsg(MarkMapReq, func() any { return new(proto.MarkMapReq) })
	c.regMsg(MarkMapRsp, func() any { return new(proto.MarkMapRsp) })
	c.regMsg(MarkNewNotify, func() any { return new(proto.MarkNewNotify) })
	c.regMsg(MarkTargetInvestigationMonsterNotify, func() any { return new(proto.MarkTargetInvestigationMonsterNotify) })
	c.regMsg(MassiveEntityElementOpBatchNotify, func() any { return new(proto.MassiveEntityElementOpBatchNotify) })
	c.regMsg(MassiveEntityStateChangedNotify, func() any { return new(proto.MassiveEntityStateChangedNotify) })
	c.regMsg(MaterialDeleteReturnNotify, func() any { return new(proto.MaterialDeleteReturnNotify) })
	c.regMsg(MaterialDeleteUpdateNotify, func() any { return new(proto.MaterialDeleteUpdateNotify) })
	c.regMsg(McoinExchangeHcoinReq, func() any { return new(proto.McoinExchangeHcoinReq) })
	c.regMsg(McoinExchangeHcoinRsp, func() any { return new(proto.McoinExchangeHcoinRsp) })
	c.regMsg(MechanicusCandidateTeamCreateReq, func() any { return new(proto.MechanicusCandidateTeamCreateReq) })
	c.regMsg(MechanicusCandidateTeamCreateRsp, func() any { return new(proto.MechanicusCandidateTeamCreateRsp) })
	c.regMsg(MechanicusCloseNotify, func() any { return new(proto.MechanicusCloseNotify) })
	c.regMsg(MechanicusCoinNotify, func() any { return new(proto.MechanicusCoinNotify) })
	c.regMsg(MechanicusLevelupGearReq, func() any { return new(proto.MechanicusLevelupGearReq) })
	c.regMsg(MechanicusLevelupGearRsp, func() any { return new(proto.MechanicusLevelupGearRsp) })
	c.regMsg(MechanicusOpenNotify, func() any { return new(proto.MechanicusOpenNotify) })
	c.regMsg(MechanicusSequenceOpenNotify, func() any { return new(proto.MechanicusSequenceOpenNotify) })
	c.regMsg(MechanicusUnlockGearReq, func() any { return new(proto.MechanicusUnlockGearReq) })
	c.regMsg(MechanicusUnlockGearRsp, func() any { return new(proto.MechanicusUnlockGearRsp) })
	c.regMsg(MeetNpcReq, func() any { return new(proto.MeetNpcReq) })
	c.regMsg(MeetNpcRsp, func() any { return new(proto.MeetNpcRsp) })
	c.regMsg(MetNpcIdListNotify, func() any { return new(proto.MetNpcIdListNotify) })
	c.regMsg(MichiaeMatsuriDarkPressureLevelUpdateNotify, func() any { return new(proto.MichiaeMatsuriDarkPressureLevelUpdateNotify) })
	c.regMsg(MichiaeMatsuriGainCrystalExpUpdateNotify, func() any { return new(proto.MichiaeMatsuriGainCrystalExpUpdateNotify) })
	c.regMsg(MichiaeMatsuriInteractStatueReq, func() any { return new(proto.MichiaeMatsuriInteractStatueReq) })
	c.regMsg(MichiaeMatsuriInteractStatueRsp, func() any { return new(proto.MichiaeMatsuriInteractStatueRsp) })
	c.regMsg(MichiaeMatsuriRemoveChallengeMarkNotify, func() any { return new(proto.MichiaeMatsuriRemoveChallengeMarkNotify) })
	c.regMsg(MichiaeMatsuriRemoveChestMarkNotify, func() any { return new(proto.MichiaeMatsuriRemoveChestMarkNotify) })
	c.regMsg(MichiaeMatsuriStartBossChallengeReq, func() any { return new(proto.MichiaeMatsuriStartBossChallengeReq) })
	c.regMsg(MichiaeMatsuriStartBossChallengeRsp, func() any { return new(proto.MichiaeMatsuriStartBossChallengeRsp) })
	c.regMsg(MichiaeMatsuriStartDarkChallengeReq, func() any { return new(proto.MichiaeMatsuriStartDarkChallengeReq) })
	c.regMsg(MichiaeMatsuriStartDarkChallengeRsp, func() any { return new(proto.MichiaeMatsuriStartDarkChallengeRsp) })
	c.regMsg(MichiaeMatsuriUnlockCrystalSkillReq, func() any { return new(proto.MichiaeMatsuriUnlockCrystalSkillReq) })
	c.regMsg(MichiaeMatsuriUnlockCrystalSkillRsp, func() any { return new(proto.MichiaeMatsuriUnlockCrystalSkillRsp) })
	c.regMsg(MiracleRingDataNotify, func() any { return new(proto.MiracleRingDataNotify) })
	c.regMsg(MiracleRingDeliverItemReq, func() any { return new(proto.MiracleRingDeliverItemReq) })
	c.regMsg(MiracleRingDeliverItemRsp, func() any { return new(proto.MiracleRingDeliverItemRsp) })
	c.regMsg(MiracleRingDestroyNotify, func() any { return new(proto.MiracleRingDestroyNotify) })
	c.regMsg(MiracleRingDropResultNotify, func() any { return new(proto.MiracleRingDropResultNotify) })
	c.regMsg(MiracleRingTakeRewardReq, func() any { return new(proto.MiracleRingTakeRewardReq) })
	c.regMsg(MiracleRingTakeRewardRsp, func() any { return new(proto.MiracleRingTakeRewardRsp) })
	c.regMsg(MistTrialFloorLevelNotify, func() any { return new(proto.MistTrialFloorLevelNotify) })
	c.regMsg(MistTrialGetChallengeMissionReq, func() any { return new(proto.MistTrialGetChallengeMissionReq) })
	c.regMsg(MistTrialGetChallengeMissionRsp, func() any { return new(proto.MistTrialGetChallengeMissionRsp) })
	c.regMsg(MistTrialGetDungeonExhibitionDataReq, func() any { return new(proto.MistTrialGetDungeonExhibitionDataReq) })
	c.regMsg(MistTrialGetDungeonExhibitionDataRsp, func() any { return new(proto.MistTrialGetDungeonExhibitionDataRsp) })
	c.regMsg(MistTrialSelectAvatarAndEnterDungeonReq, func() any { return new(proto.MistTrialSelectAvatarAndEnterDungeonReq) })
	c.regMsg(MistTrialSelectAvatarAndEnterDungeonRsp, func() any { return new(proto.MistTrialSelectAvatarAndEnterDungeonRsp) })
	c.regMsg(MistTrialSettleNotify, func() any { return new(proto.MistTrialSettleNotify) })
	c.regMsg(MonsterAIConfigHashNotify, func() any { return new(proto.MonsterAIConfigHashNotify) })
	c.regMsg(MonsterAlertChangeNotify, func() any { return new(proto.MonsterAlertChangeNotify) })
	c.regMsg(MonsterForceAlertNotify, func() any { return new(proto.MonsterForceAlertNotify) })
	c.regMsg(MonsterPointArrayRouteUpdateNotify, func() any { return new(proto.MonsterPointArrayRouteUpdateNotify) })
	c.regMsg(MonsterSummonTagNotify, func() any { return new(proto.MonsterSummonTagNotify) })
	c.regMsg(MpBlockNotify, func() any { return new(proto.MpBlockNotify) })
	c.regMsg(MpPlayGuestReplyInviteReq, func() any { return new(proto.MpPlayGuestReplyInviteReq) })
	c.regMsg(MpPlayGuestReplyInviteRsp, func() any { return new(proto.MpPlayGuestReplyInviteRsp) })
	c.regMsg(MpPlayGuestReplyNotify, func() any { return new(proto.MpPlayGuestReplyNotify) })
	c.regMsg(MpPlayInviteResultNotify, func() any { return new(proto.MpPlayInviteResultNotify) })
	c.regMsg(MpPlayOwnerCheckReq, func() any { return new(proto.MpPlayOwnerCheckReq) })
	c.regMsg(MpPlayOwnerCheckRsp, func() any { return new(proto.MpPlayOwnerCheckRsp) })
	c.regMsg(MpPlayOwnerInviteNotify, func() any { return new(proto.MpPlayOwnerInviteNotify) })
	c.regMsg(MpPlayOwnerStartInviteReq, func() any { return new(proto.MpPlayOwnerStartInviteReq) })
	c.regMsg(MpPlayOwnerStartInviteRsp, func() any { return new(proto.MpPlayOwnerStartInviteRsp) })
	c.regMsg(MpPlayPrepareInterruptNotify, func() any { return new(proto.MpPlayPrepareInterruptNotify) })
	c.regMsg(MpPlayPrepareNotify, func() any { return new(proto.MpPlayPrepareNotify) })
	c.regMsg(MultistagePlayEndNotify, func() any { return new(proto.MultistagePlayEndNotify) })
	c.regMsg(MultistagePlayFinishStageReq, func() any { return new(proto.MultistagePlayFinishStageReq) })
	c.regMsg(MultistagePlayFinishStageRsp, func() any { return new(proto.MultistagePlayFinishStageRsp) })
	c.regMsg(MultistagePlayInfoNotify, func() any { return new(proto.MultistagePlayInfoNotify) })
	c.regMsg(MultistagePlaySettleNotify, func() any { return new(proto.MultistagePlaySettleNotify) })
	c.regMsg(MultistagePlayStageEndNotify, func() any { return new(proto.MultistagePlayStageEndNotify) })
	c.regMsg(MuqadasPotionActivityEnterDungeonReq, func() any { return new(proto.MuqadasPotionActivityEnterDungeonReq) })
	c.regMsg(MuqadasPotionActivityEnterDungeonRsp, func() any { return new(proto.MuqadasPotionActivityEnterDungeonRsp) })
	c.regMsg(MuqadasPotionCaptureWeaknessReq, func() any { return new(proto.MuqadasPotionCaptureWeaknessReq) })
	c.regMsg(MuqadasPotionCaptureWeaknessRsp, func() any { return new(proto.MuqadasPotionCaptureWeaknessRsp) })
	c.regMsg(MuqadasPotionDungeonSettleNotify, func() any { return new(proto.MuqadasPotionDungeonSettleNotify) })
	c.regMsg(MuqadasPotionRestartDungeonReq, func() any { return new(proto.MuqadasPotionRestartDungeonReq) })
	c.regMsg(MuqadasPotionRestartDungeonRsp, func() any { return new(proto.MuqadasPotionRestartDungeonRsp) })
	c.regMsg(MusicGameSettleReq, func() any { return new(proto.MusicGameSettleReq) })
	c.regMsg(MusicGameSettleRsp, func() any { return new(proto.MusicGameSettleRsp) })
	c.regMsg(MusicGameStartReq, func() any { return new(proto.MusicGameStartReq) })
	c.regMsg(MusicGameStartRsp, func() any { return new(proto.MusicGameStartRsp) })
	c.regMsg(NavMeshStatsNotify, func() any { return new(proto.NavMeshStatsNotify) })
	c.regMsg(NicknameAuditConfigNotify, func() any { return new(proto.NicknameAuditConfigNotify) })
	c.regMsg(NightCrowGadgetObservationMatchReq, func() any { return new(proto.NightCrowGadgetObservationMatchReq) })
	c.regMsg(NightCrowGadgetObservationMatchRsp, func() any { return new(proto.NightCrowGadgetObservationMatchRsp) })
	c.regMsg(NormalUidOpNotify, func() any { return new(proto.NormalUidOpNotify) })
	c.regMsg(NpcTalkReq, func() any { return new(proto.NpcTalkReq) })
	c.regMsg(NpcTalkRsp, func() any { return new(proto.NpcTalkRsp) })
	c.regMsg(NpcTalkStateNotify, func() any { return new(proto.NpcTalkStateNotify) })
	c.regMsg(ObstacleModifyNotify, func() any { return new(proto.ObstacleModifyNotify) })
	c.regMsg(OfferingInteractReq, func() any { return new(proto.OfferingInteractReq) })
	c.regMsg(OfferingInteractRsp, func() any { return new(proto.OfferingInteractRsp) })
	c.regMsg(OpActivityDataNotify, func() any { return new(proto.OpActivityDataNotify) })
	c.regMsg(OpActivityStateNotify, func() any { return new(proto.OpActivityStateNotify) })
	c.regMsg(OpActivityUpdateNotify, func() any { return new(proto.OpActivityUpdateNotify) })
	c.regMsg(OpenBlossomCircleCampGuideNotify, func() any { return new(proto.OpenBlossomCircleCampGuideNotify) })
	c.regMsg(OpenStateChangeNotify, func() any { return new(proto.OpenStateChangeNotify) })
	c.regMsg(OpenStateUpdateNotify, func() any { return new(proto.OpenStateUpdateNotify) })
	c.regMsg(OrderDisplayNotify, func() any { return new(proto.OrderDisplayNotify) })
	c.regMsg(OrderFinishNotify, func() any { return new(proto.OrderFinishNotify) })
	c.regMsg(OtherPlayerEnterHomeNotify, func() any { return new(proto.OtherPlayerEnterHomeNotify) })
	c.regMsg(OutStuckCustomDungeonReq, func() any { return new(proto.OutStuckCustomDungeonReq) })
	c.regMsg(OutStuckCustomDungeonRsp, func() any { return new(proto.OutStuckCustomDungeonRsp) })
	c.regMsg(PSNBlackListNotify, func() any { return new(proto.PSNBlackListNotify) })
	c.regMsg(PSNFriendListNotify, func() any { return new(proto.PSNFriendListNotify) })
	c.regMsg(PSPlayerApplyEnterMpReq, func() any { return new(proto.PSPlayerApplyEnterMpReq) })
	c.regMsg(PSPlayerApplyEnterMpRsp, func() any { return new(proto.PSPlayerApplyEnterMpRsp) })
	c.regMsg(ParentQuestInferenceDataNotify, func() any { return new(proto.ParentQuestInferenceDataNotify) })
	c.regMsg(PathfindingEnterSceneReq, func() any { return new(proto.PathfindingEnterSceneReq) })
	c.regMsg(PathfindingEnterSceneRsp, func() any { return new(proto.PathfindingEnterSceneRsp) })
	c.regMsg(PathfindingPingNotify, func() any { return new(proto.PathfindingPingNotify) })
	c.regMsg(PersistentDungeonSwitchAvatarReq, func() any { return new(proto.PersistentDungeonSwitchAvatarReq) })
	c.regMsg(PersistentDungeonSwitchAvatarRsp, func() any { return new(proto.PersistentDungeonSwitchAvatarRsp) })
	c.regMsg(PersonalLineAllDataReq, func() any { return new(proto.PersonalLineAllDataReq) })
	c.regMsg(PersonalLineAllDataRsp, func() any { return new(proto.PersonalLineAllDataRsp) })
	c.regMsg(PersonalLineNewUnlockNotify, func() any { return new(proto.PersonalLineNewUnlockNotify) })
	c.regMsg(PersonalSceneJumpReq, func() any { return new(proto.PersonalSceneJumpReq) })
	c.regMsg(PersonalSceneJumpRsp, func() any { return new(proto.PersonalSceneJumpRsp) })
	c.regMsg(PhotoActivityClientViewReq, func() any { return new(proto.PhotoActivityClientViewReq) })
	c.regMsg(PhotoActivityClientViewRsp, func() any { return new(proto.PhotoActivityClientViewRsp) })
	c.regMsg(PhotoActivityFinishReq, func() any { return new(proto.PhotoActivityFinishReq) })
	c.regMsg(PhotoActivityFinishRsp, func() any { return new(proto.PhotoActivityFinishRsp) })
	c.regMsg(PingReq, func() any { return new(proto.PingReq) })
	c.regMsg(PingRsp, func() any { return new(proto.PingRsp) })
	c.regMsg(PlantFlowerAcceptAllGiveFlowerReq, func() any { return new(proto.PlantFlowerAcceptAllGiveFlowerReq) })
	c.regMsg(PlantFlowerAcceptAllGiveFlowerRsp, func() any { return new(proto.PlantFlowerAcceptAllGiveFlowerRsp) })
	c.regMsg(PlantFlowerAcceptGiveFlowerReq, func() any { return new(proto.PlantFlowerAcceptGiveFlowerReq) })
	c.regMsg(PlantFlowerAcceptGiveFlowerRsp, func() any { return new(proto.PlantFlowerAcceptGiveFlowerRsp) })
	c.regMsg(PlantFlowerEditFlowerCombinationReq, func() any { return new(proto.PlantFlowerEditFlowerCombinationReq) })
	c.regMsg(PlantFlowerEditFlowerCombinationRsp, func() any { return new(proto.PlantFlowerEditFlowerCombinationRsp) })
	c.regMsg(PlantFlowerGetCanGiveFriendFlowerReq, func() any { return new(proto.PlantFlowerGetCanGiveFriendFlowerReq) })
	c.regMsg(PlantFlowerGetCanGiveFriendFlowerRsp, func() any { return new(proto.PlantFlowerGetCanGiveFriendFlowerRsp) })
	c.regMsg(PlantFlowerGetFriendFlowerWishListReq, func() any { return new(proto.PlantFlowerGetFriendFlowerWishListReq) })
	c.regMsg(PlantFlowerGetFriendFlowerWishListRsp, func() any { return new(proto.PlantFlowerGetFriendFlowerWishListRsp) })
	c.regMsg(PlantFlowerGetRecvFlowerListReq, func() any { return new(proto.PlantFlowerGetRecvFlowerListReq) })
	c.regMsg(PlantFlowerGetRecvFlowerListRsp, func() any { return new(proto.PlantFlowerGetRecvFlowerListRsp) })
	c.regMsg(PlantFlowerGetSeedInfoReq, func() any { return new(proto.PlantFlowerGetSeedInfoReq) })
	c.regMsg(PlantFlowerGetSeedInfoRsp, func() any { return new(proto.PlantFlowerGetSeedInfoRsp) })
	c.regMsg(PlantFlowerGiveFriendFlowerReq, func() any { return new(proto.PlantFlowerGiveFriendFlowerReq) })
	c.regMsg(PlantFlowerGiveFriendFlowerRsp, func() any { return new(proto.PlantFlowerGiveFriendFlowerRsp) })
	c.regMsg(PlantFlowerHaveRecvFlowerNotify, func() any { return new(proto.PlantFlowerHaveRecvFlowerNotify) })
	c.regMsg(PlantFlowerSetFlowerWishReq, func() any { return new(proto.PlantFlowerSetFlowerWishReq) })
	c.regMsg(PlantFlowerSetFlowerWishRsp, func() any { return new(proto.PlantFlowerSetFlowerWishRsp) })
	c.regMsg(PlantFlowerTakeSeedRewardReq, func() any { return new(proto.PlantFlowerTakeSeedRewardReq) })
	c.regMsg(PlantFlowerTakeSeedRewardRsp, func() any { return new(proto.PlantFlowerTakeSeedRewardRsp) })
	c.regMsg(PlatformChangeRouteNotify, func() any { return new(proto.PlatformChangeRouteNotify) })
	c.regMsg(PlatformStartRouteNotify, func() any { return new(proto.PlatformStartRouteNotify) })
	c.regMsg(PlatformStopRouteNotify, func() any { return new(proto.PlatformStopRouteNotify) })
	c.regMsg(PlayerAllowEnterMpAfterAgreeMatchNotify, func() any { return new(proto.PlayerAllowEnterMpAfterAgreeMatchNotify) })
	c.regMsg(PlayerApplyEnterHomeNotify, func() any { return new(proto.PlayerApplyEnterHomeNotify) })
	c.regMsg(PlayerApplyEnterHomeResultNotify, func() any { return new(proto.PlayerApplyEnterHomeResultNotify) })
	c.regMsg(PlayerApplyEnterHomeResultReq, func() any { return new(proto.PlayerApplyEnterHomeResultReq) })
	c.regMsg(PlayerApplyEnterHomeResultRsp, func() any { return new(proto.PlayerApplyEnterHomeResultRsp) })
	c.regMsg(PlayerApplyEnterMpAfterMatchAgreedNotify, func() any { return new(proto.PlayerApplyEnterMpAfterMatchAgreedNotify) })
	c.regMsg(PlayerApplyEnterMpNotify, func() any { return new(proto.PlayerApplyEnterMpNotify) })
	c.regMsg(PlayerApplyEnterMpReq, func() any { return new(proto.PlayerApplyEnterMpReq) })
	c.regMsg(PlayerApplyEnterMpResultNotify, func() any { return new(proto.PlayerApplyEnterMpResultNotify) })
	c.regMsg(PlayerApplyEnterMpResultReq, func() any { return new(proto.PlayerApplyEnterMpResultReq) })
	c.regMsg(PlayerApplyEnterMpResultRsp, func() any { return new(proto.PlayerApplyEnterMpResultRsp) })
	c.regMsg(PlayerApplyEnterMpRsp, func() any { return new(proto.PlayerApplyEnterMpRsp) })
	c.regMsg(PlayerCancelMatchReq, func() any { return new(proto.PlayerCancelMatchReq) })
	c.regMsg(PlayerCancelMatchRsp, func() any { return new(proto.PlayerCancelMatchRsp) })
	c.regMsg(PlayerChatCDNotify, func() any { return new(proto.PlayerChatCDNotify) })
	c.regMsg(PlayerChatNotify, func() any { return new(proto.PlayerChatNotify) })
	c.regMsg(PlayerChatReq, func() any { return new(proto.PlayerChatReq) })
	c.regMsg(PlayerChatRsp, func() any { return new(proto.PlayerChatRsp) })
	c.regMsg(PlayerCompoundMaterialBoostReq, func() any { return new(proto.PlayerCompoundMaterialBoostReq) })
	c.regMsg(PlayerCompoundMaterialBoostRsp, func() any { return new(proto.PlayerCompoundMaterialBoostRsp) })
	c.regMsg(PlayerCompoundMaterialReq, func() any { return new(proto.PlayerCompoundMaterialReq) })
	c.regMsg(PlayerCompoundMaterialRsp, func() any { return new(proto.PlayerCompoundMaterialRsp) })
	c.regMsg(PlayerConfirmMatchReq, func() any { return new(proto.PlayerConfirmMatchReq) })
	c.regMsg(PlayerConfirmMatchRsp, func() any { return new(proto.PlayerConfirmMatchRsp) })
	c.regMsg(PlayerCookArgsReq, func() any { return new(proto.PlayerCookArgsReq) })
	c.regMsg(PlayerCookArgsRsp, func() any { return new(proto.PlayerCookArgsRsp) })
	c.regMsg(PlayerCookReq, func() any { return new(proto.PlayerCookReq) })
	c.regMsg(PlayerCookRsp, func() any { return new(proto.PlayerCookRsp) })
	c.regMsg(PlayerDataNotify, func() any { return new(proto.PlayerDataNotify) })
	c.regMsg(PlayerDeathZoneNotify, func() any { return new(proto.PlayerDeathZoneNotify) })
	c.regMsg(PlayerEnterDungeonReq, func() any { return new(proto.PlayerEnterDungeonReq) })
	c.regMsg(PlayerEnterDungeonRsp, func() any { return new(proto.PlayerEnterDungeonRsp) })
	c.regMsg(PlayerEnterSceneInfoNotify, func() any { return new(proto.PlayerEnterSceneInfoNotify) })
	c.regMsg(PlayerEnterSceneNotify, func() any { return new(proto.PlayerEnterSceneNotify) })
	c.regMsg(PlayerEyePointStateNotify, func() any { return new(proto.PlayerEyePointStateNotify) })
	c.regMsg(PlayerFishingDataNotify, func() any { return new(proto.PlayerFishingDataNotify) })
	c.regMsg(PlayerForceExitReq, func() any { return new(proto.PlayerForceExitReq) })
	c.regMsg(PlayerForceExitRsp, func() any { return new(proto.PlayerForceExitRsp) })
	c.regMsg(PlayerGCGMatchConfirmNotify, func() any { return new(proto.PlayerGCGMatchConfirmNotify) })
	c.regMsg(PlayerGCGMatchDismissNotify, func() any { return new(proto.PlayerGCGMatchDismissNotify) })
	c.regMsg(PlayerGameTimeNotify, func() any { return new(proto.PlayerGameTimeNotify) })
	c.regMsg(PlayerGeneralMatchConfirmNotify, func() any { return new(proto.PlayerGeneralMatchConfirmNotify) })
	c.regMsg(PlayerGeneralMatchDismissNotify, func() any { return new(proto.PlayerGeneralMatchDismissNotify) })
	c.regMsg(PlayerGetForceQuitBanInfoReq, func() any { return new(proto.PlayerGetForceQuitBanInfoReq) })
	c.regMsg(PlayerGetForceQuitBanInfoRsp, func() any { return new(proto.PlayerGetForceQuitBanInfoRsp) })
	c.regMsg(PlayerHomeCompInfoNotify, func() any { return new(proto.PlayerHomeCompInfoNotify) })
	c.regMsg(PlayerInjectFixNotify, func() any { return new(proto.PlayerInjectFixNotify) })
	c.regMsg(PlayerInvestigationAllInfoNotify, func() any { return new(proto.PlayerInvestigationAllInfoNotify) })
	c.regMsg(PlayerInvestigationNotify, func() any { return new(proto.PlayerInvestigationNotify) })
	c.regMsg(PlayerInvestigationTargetNotify, func() any { return new(proto.PlayerInvestigationTargetNotify) })
	c.regMsg(PlayerLevelRewardUpdateNotify, func() any { return new(proto.PlayerLevelRewardUpdateNotify) })
	c.regMsg(PlayerLoginReq, func() any { return new(proto.PlayerLoginReq) })
	c.regMsg(PlayerLoginRsp, func() any { return new(proto.PlayerLoginRsp) })
	c.regMsg(PlayerLogoutNotify, func() any { return new(proto.PlayerLogoutNotify) })
	c.regMsg(PlayerLogoutReq, func() any { return new(proto.PlayerLogoutReq) })
	c.regMsg(PlayerLogoutRsp, func() any { return new(proto.PlayerLogoutRsp) })
	c.regMsg(PlayerLuaShellNotify, func() any { return new(proto.PlayerLuaShellNotify) })
	c.regMsg(PlayerMatchAgreedResultNotify, func() any { return new(proto.PlayerMatchAgreedResultNotify) })
	c.regMsg(PlayerMatchInfoNotify, func() any { return new(proto.PlayerMatchInfoNotify) })
	c.regMsg(PlayerMatchStopNotify, func() any { return new(proto.PlayerMatchStopNotify) })
	c.regMsg(PlayerMatchSuccNotify, func() any { return new(proto.PlayerMatchSuccNotify) })
	c.regMsg(PlayerNicknameAuditDataNotify, func() any { return new(proto.PlayerNicknameAuditDataNotify) })
	c.regMsg(PlayerNicknameNotify, func() any { return new(proto.PlayerNicknameNotify) })
	c.regMsg(PlayerOfferingDataNotify, func() any { return new(proto.PlayerOfferingDataNotify) })
	c.regMsg(PlayerOfferingReq, func() any { return new(proto.PlayerOfferingReq) })
	c.regMsg(PlayerOfferingRsp, func() any { return new(proto.PlayerOfferingRsp) })
	c.regMsg(PlayerPreEnterMpNotify, func() any { return new(proto.PlayerPreEnterMpNotify) })
	c.regMsg(PlayerPropChangeNotify, func() any { return new(proto.PlayerPropChangeNotify) })
	c.regMsg(PlayerPropChangeReasonNotify, func() any { return new(proto.PlayerPropChangeReasonNotify) })
	c.regMsg(PlayerPropNotify, func() any { return new(proto.PlayerPropNotify) })
	c.regMsg(PlayerQuitDungeonReq, func() any { return new(proto.PlayerQuitDungeonReq) })
	c.regMsg(PlayerQuitDungeonRsp, func() any { return new(proto.PlayerQuitDungeonRsp) })
	c.regMsg(PlayerQuitFromHomeNotify, func() any { return new(proto.PlayerQuitFromHomeNotify) })
	c.regMsg(PlayerQuitFromMpNotify, func() any { return new(proto.PlayerQuitFromMpNotify) })
	c.regMsg(PlayerRandomCookReq, func() any { return new(proto.PlayerRandomCookReq) })
	c.regMsg(PlayerRandomCookRsp, func() any { return new(proto.PlayerRandomCookRsp) })
	c.regMsg(PlayerRechargeDataNotify, func() any { return new(proto.PlayerRechargeDataNotify) })
	c.regMsg(PlayerReportReq, func() any { return new(proto.PlayerReportReq) })
	c.regMsg(PlayerReportRsp, func() any { return new(proto.PlayerReportRsp) })
	c.regMsg(PlayerRoutineDataNotify, func() any { return new(proto.PlayerRoutineDataNotify) })
	c.regMsg(PlayerSetLanguageReq, func() any { return new(proto.PlayerSetLanguageReq) })
	c.regMsg(PlayerSetLanguageRsp, func() any { return new(proto.PlayerSetLanguageRsp) })
	c.regMsg(PlayerSetOnlyMPWithPSPlayerReq, func() any { return new(proto.PlayerSetOnlyMPWithPSPlayerReq) })
	c.regMsg(PlayerSetOnlyMPWithPSPlayerRsp, func() any { return new(proto.PlayerSetOnlyMPWithPSPlayerRsp) })
	c.regMsg(PlayerSetPauseReq, func() any { return new(proto.PlayerSetPauseReq) })
	c.regMsg(PlayerSetPauseRsp, func() any { return new(proto.PlayerSetPauseRsp) })
	c.regMsg(PlayerSignatureAuditDataNotify, func() any { return new(proto.PlayerSignatureAuditDataNotify) })
	c.regMsg(PlayerSignatureNotify, func() any { return new(proto.PlayerSignatureNotify) })
	c.regMsg(PlayerStartMatchReq, func() any { return new(proto.PlayerStartMatchReq) })
	c.regMsg(PlayerStartMatchRsp, func() any { return new(proto.PlayerStartMatchRsp) })
	c.regMsg(PlayerStoreNotify, func() any { return new(proto.PlayerStoreNotify) })
	c.regMsg(PlayerTimeNotify, func() any { return new(proto.PlayerTimeNotify) })
	c.regMsg(PlayerWorldSceneInfoListNotify, func() any { return new(proto.PlayerWorldSceneInfoListNotify) })
	c.regMsg(PostEnterSceneReq, func() any { return new(proto.PostEnterSceneReq) })
	c.regMsg(PostEnterSceneRsp, func() any { return new(proto.PostEnterSceneRsp) })
	c.regMsg(PotionEnterDungeonNotify, func() any { return new(proto.PotionEnterDungeonNotify) })
	c.regMsg(PotionEnterDungeonReq, func() any { return new(proto.PotionEnterDungeonReq) })
	c.regMsg(PotionEnterDungeonRsp, func() any { return new(proto.PotionEnterDungeonRsp) })
	c.regMsg(PotionResetChallengeReq, func() any { return new(proto.PotionResetChallengeReq) })
	c.regMsg(PotionResetChallengeRsp, func() any { return new(proto.PotionResetChallengeRsp) })
	c.regMsg(PotionRestartDungeonReq, func() any { return new(proto.PotionRestartDungeonReq) })
	c.regMsg(PotionRestartDungeonRsp, func() any { return new(proto.PotionRestartDungeonRsp) })
	c.regMsg(PotionSaveDungeonResultReq, func() any { return new(proto.PotionSaveDungeonResultReq) })
	c.regMsg(PotionSaveDungeonResultRsp, func() any { return new(proto.PotionSaveDungeonResultRsp) })
	c.regMsg(PrivateChatNotify, func() any { return new(proto.PrivateChatNotify) })
	c.regMsg(PrivateChatReq, func() any { return new(proto.PrivateChatReq) })
	c.regMsg(PrivateChatRsp, func() any { return new(proto.PrivateChatRsp) })
	c.regMsg(ProfilePictureChangeNotify, func() any { return new(proto.ProfilePictureChangeNotify) })
	c.regMsg(ProjectorOptionReq, func() any { return new(proto.ProjectorOptionReq) })
	c.regMsg(ProjectorOptionRsp, func() any { return new(proto.ProjectorOptionRsp) })
	c.regMsg(ProudSkillChangeNotify, func() any { return new(proto.ProudSkillChangeNotify) })
	c.regMsg(ProudSkillExtraLevelNotify, func() any { return new(proto.ProudSkillExtraLevelNotify) })
	c.regMsg(ProudSkillUpgradeReq, func() any { return new(proto.ProudSkillUpgradeReq) })
	c.regMsg(ProudSkillUpgradeRsp, func() any { return new(proto.ProudSkillUpgradeRsp) })
	c.regMsg(PublishCustomDungeonReq, func() any { return new(proto.PublishCustomDungeonReq) })
	c.regMsg(PublishCustomDungeonRsp, func() any { return new(proto.PublishCustomDungeonRsp) })
	c.regMsg(PublishUgcReq, func() any { return new(proto.PublishUgcReq) })
	c.regMsg(PublishUgcRsp, func() any { return new(proto.PublishUgcRsp) })
	c.regMsg(PullPrivateChatReq, func() any { return new(proto.PullPrivateChatReq) })
	c.regMsg(PullPrivateChatRsp, func() any { return new(proto.PullPrivateChatRsp) })
	c.regMsg(PullRecentChatReq, func() any { return new(proto.PullRecentChatReq) })
	c.regMsg(PullRecentChatRsp, func() any { return new(proto.PullRecentChatRsp) })
	c.regMsg(PushTipsAllDataNotify, func() any { return new(proto.PushTipsAllDataNotify) })
	c.regMsg(PushTipsChangeNotify, func() any { return new(proto.PushTipsChangeNotify) })
	c.regMsg(PushTipsReadFinishReq, func() any { return new(proto.PushTipsReadFinishReq) })
	c.regMsg(PushTipsReadFinishRsp, func() any { return new(proto.PushTipsReadFinishRsp) })
	c.regMsg(QueryCodexMonsterBeKilledNumReq, func() any { return new(proto.QueryCodexMonsterBeKilledNumReq) })
	c.regMsg(QueryCodexMonsterBeKilledNumRsp, func() any { return new(proto.QueryCodexMonsterBeKilledNumRsp) })
	c.regMsg(QueryPathReq, func() any { return new(proto.QueryPathReq) })
	c.regMsg(QueryPathRsp, func() any { return new(proto.QueryPathRsp) })
	c.regMsg(QuestCreateEntityReq, func() any { return new(proto.QuestCreateEntityReq) })
	c.regMsg(QuestCreateEntityRsp, func() any { return new(proto.QuestCreateEntityRsp) })
	c.regMsg(QuestDelNotify, func() any { return new(proto.QuestDelNotify) })
	c.regMsg(QuestDestroyEntityReq, func() any { return new(proto.QuestDestroyEntityReq) })
	c.regMsg(QuestDestroyEntityRsp, func() any { return new(proto.QuestDestroyEntityRsp) })
	c.regMsg(QuestDestroyNpcReq, func() any { return new(proto.QuestDestroyNpcReq) })
	c.regMsg(QuestDestroyNpcRsp, func() any { return new(proto.QuestDestroyNpcRsp) })
	c.regMsg(QuestGlobalVarNotify, func() any { return new(proto.QuestGlobalVarNotify) })
	c.regMsg(QuestListNotify, func() any { return new(proto.QuestListNotify) })
	c.regMsg(QuestListUpdateNotify, func() any { return new(proto.QuestListUpdateNotify) })
	c.regMsg(QuestProgressUpdateNotify, func() any { return new(proto.QuestProgressUpdateNotify) })
	c.regMsg(QuestTransmitReq, func() any { return new(proto.QuestTransmitReq) })
	c.regMsg(QuestTransmitRsp, func() any { return new(proto.QuestTransmitRsp) })
	c.regMsg(QuestUpdateQuestTimeVarNotify, func() any { return new(proto.QuestUpdateQuestTimeVarNotify) })
	c.regMsg(QuestUpdateQuestVarNotify, func() any { return new(proto.QuestUpdateQuestVarNotify) })
	c.regMsg(QuestUpdateQuestVarReq, func() any { return new(proto.QuestUpdateQuestVarReq) })
	c.regMsg(QuestUpdateQuestVarRsp, func() any { return new(proto.QuestUpdateQuestVarRsp) })
	c.regMsg(QuickOpenActivityReq, func() any { return new(proto.QuickOpenActivityReq) })
	c.regMsg(QuickOpenActivityRsp, func() any { return new(proto.QuickOpenActivityRsp) })
	c.regMsg(QuickUseWidgetReq, func() any { return new(proto.QuickUseWidgetReq) })
	c.regMsg(QuickUseWidgetRsp, func() any { return new(proto.QuickUseWidgetRsp) })
	c.regMsg(ReadMailNotify, func() any { return new(proto.ReadMailNotify) })
	c.regMsg(ReadNicknameAuditReq, func() any { return new(proto.ReadNicknameAuditReq) })
	c.regMsg(ReadNicknameAuditRsp, func() any { return new(proto.ReadNicknameAuditRsp) })
	c.regMsg(ReadPrivateChatReq, func() any { return new(proto.ReadPrivateChatReq) })
	c.regMsg(ReadPrivateChatRsp, func() any { return new(proto.ReadPrivateChatRsp) })
	c.regMsg(ReadSignatureAuditReq, func() any { return new(proto.ReadSignatureAuditReq) })
	c.regMsg(ReadSignatureAuditRsp, func() any { return new(proto.ReadSignatureAuditRsp) })
	c.regMsg(ReceivedTrialAvatarActivityRewardReq, func() any { return new(proto.ReceivedTrialAvatarActivityRewardReq) })
	c.regMsg(ReceivedTrialAvatarActivityRewardRsp, func() any { return new(proto.ReceivedTrialAvatarActivityRewardRsp) })
	c.regMsg(RechargeReq, func() any { return new(proto.RechargeReq) })
	c.regMsg(RechargeRsp, func() any { return new(proto.RechargeRsp) })
	c.regMsg(RedeemLegendaryKeyReq, func() any { return new(proto.RedeemLegendaryKeyReq) })
	c.regMsg(RedeemLegendaryKeyRsp, func() any { return new(proto.RedeemLegendaryKeyRsp) })
	c.regMsg(ReformFireworksReq, func() any { return new(proto.ReformFireworksReq) })
	c.regMsg(ReformFireworksRsp, func() any { return new(proto.ReformFireworksRsp) })
	c.regMsg(RefreshBackgroundAvatarReq, func() any { return new(proto.RefreshBackgroundAvatarReq) })
	c.regMsg(RefreshBackgroundAvatarRsp, func() any { return new(proto.RefreshBackgroundAvatarRsp) })
	c.regMsg(RefreshEntityAuthNotify, func() any { return new(proto.RefreshEntityAuthNotify) })
	c.regMsg(RefreshRogueDiaryCardReq, func() any { return new(proto.RefreshRogueDiaryCardReq) })
	c.regMsg(RefreshRogueDiaryCardRsp, func() any { return new(proto.RefreshRogueDiaryCardRsp) })
	c.regMsg(RefreshRoguelikeDungeonCardReq, func() any { return new(proto.RefreshRoguelikeDungeonCardReq) })
	c.regMsg(RefreshRoguelikeDungeonCardRsp, func() any { return new(proto.RefreshRoguelikeDungeonCardRsp) })
	c.regMsg(RegionSearchChangeRegionNotify, func() any { return new(proto.RegionSearchChangeRegionNotify) })
	c.regMsg(RegionSearchNotify, func() any { return new(proto.RegionSearchNotify) })
	c.regMsg(RegionalPlayInfoNotify, func() any { return new(proto.RegionalPlayInfoNotify) })
	c.regMsg(ReliquaryDecomposeReq, func() any { return new(proto.ReliquaryDecomposeReq) })
	c.regMsg(ReliquaryDecomposeRsp, func() any { return new(proto.ReliquaryDecomposeRsp) })
	c.regMsg(ReliquaryPromoteReq, func() any { return new(proto.ReliquaryPromoteReq) })
	c.regMsg(ReliquaryPromoteRsp, func() any { return new(proto.ReliquaryPromoteRsp) })
	c.regMsg(ReliquaryUpgradeReq, func() any { return new(proto.ReliquaryUpgradeReq) })
	c.regMsg(ReliquaryUpgradeRsp, func() any { return new(proto.ReliquaryUpgradeRsp) })
	c.regMsg(RemotePlayerWidgetNotify, func() any { return new(proto.RemotePlayerWidgetNotify) })
	c.regMsg(RemoveBlacklistReq, func() any { return new(proto.RemoveBlacklistReq) })
	c.regMsg(RemoveBlacklistRsp, func() any { return new(proto.RemoveBlacklistRsp) })
	c.regMsg(RemoveCustomDungeonReq, func() any { return new(proto.RemoveCustomDungeonReq) })
	c.regMsg(RemoveCustomDungeonRsp, func() any { return new(proto.RemoveCustomDungeonRsp) })
	c.regMsg(RemoveRandTaskInfoNotify, func() any { return new(proto.RemoveRandTaskInfoNotify) })
	c.regMsg(ReplayCustomDungeonReq, func() any { return new(proto.ReplayCustomDungeonReq) })
	c.regMsg(ReplayCustomDungeonRsp, func() any { return new(proto.ReplayCustomDungeonRsp) })
	c.regMsg(ReportFightAntiCheatNotify, func() any { return new(proto.ReportFightAntiCheatNotify) })
	c.regMsg(ReportTrackingIOInfoNotify, func() any { return new(proto.ReportTrackingIOInfoNotify) })
	c.regMsg(RequestLiveInfoReq, func() any { return new(proto.RequestLiveInfoReq) })
	c.regMsg(RequestLiveInfoRsp, func() any { return new(proto.RequestLiveInfoRsp) })
	c.regMsg(ReserveRogueDiaryAvatarReq, func() any { return new(proto.ReserveRogueDiaryAvatarReq) })
	c.regMsg(ReserveRogueDiaryAvatarRsp, func() any { return new(proto.ReserveRogueDiaryAvatarRsp) })
	c.regMsg(ResetRogueDiaryPlayReq, func() any { return new(proto.ResetRogueDiaryPlayReq) })
	c.regMsg(ResetRogueDiaryPlayRsp, func() any { return new(proto.ResetRogueDiaryPlayRsp) })
	c.regMsg(ResinCardDataUpdateNotify, func() any { return new(proto.ResinCardDataUpdateNotify) })
	c.regMsg(ResinChangeNotify, func() any { return new(proto.ResinChangeNotify) })
	c.regMsg(RestartEffigyChallengeReq, func() any { return new(proto.RestartEffigyChallengeReq) })
	c.regMsg(RestartEffigyChallengeRsp, func() any { return new(proto.RestartEffigyChallengeRsp) })
	c.regMsg(ResumeRogueDiaryDungeonReq, func() any { return new(proto.ResumeRogueDiaryDungeonReq) })
	c.regMsg(ResumeRogueDiaryDungeonRsp, func() any { return new(proto.ResumeRogueDiaryDungeonRsp) })
	c.regMsg(RetryCurRogueDiaryDungeonReq, func() any { return new(proto.RetryCurRogueDiaryDungeonReq) })
	c.regMsg(RetryCurRogueDiaryDungeonRsp, func() any { return new(proto.RetryCurRogueDiaryDungeonRsp) })
	c.regMsg(ReunionActivateNotify, func() any { return new(proto.ReunionActivateNotify) })
	c.regMsg(ReunionBriefInfoReq, func() any { return new(proto.ReunionBriefInfoReq) })
	c.regMsg(ReunionBriefInfoRsp, func() any { return new(proto.ReunionBriefInfoRsp) })
	c.regMsg(ReunionDailyRefreshNotify, func() any { return new(proto.ReunionDailyRefreshNotify) })
	c.regMsg(ReunionPrivilegeChangeNotify, func() any { return new(proto.ReunionPrivilegeChangeNotify) })
	c.regMsg(ReunionSettleNotify, func() any { return new(proto.ReunionSettleNotify) })
	c.regMsg(RobotPushPlayerDataNotify, func() any { return new(proto.RobotPushPlayerDataNotify) })
	c.regMsg(RogueCellUpdateNotify, func() any { return new(proto.RogueCellUpdateNotify) })
	c.regMsg(RogueDiaryCoinAddNotify, func() any { return new(proto.RogueDiaryCoinAddNotify) })
	c.regMsg(RogueDiaryDungeonInfoNotify, func() any { return new(proto.RogueDiaryDungeonInfoNotify) })
	c.regMsg(RogueDiaryDungeonSettleNotify, func() any { return new(proto.RogueDiaryDungeonSettleNotify) })
	c.regMsg(RogueDiaryRepairInfoNotify, func() any { return new(proto.RogueDiaryRepairInfoNotify) })
	c.regMsg(RogueDiaryReviveAvatarReq, func() any { return new(proto.RogueDiaryReviveAvatarReq) })
	c.regMsg(RogueDiaryReviveAvatarRsp, func() any { return new(proto.RogueDiaryReviveAvatarRsp) })
	c.regMsg(RogueDiaryTiredAvatarNotify, func() any { return new(proto.RogueDiaryTiredAvatarNotify) })
	c.regMsg(RogueDungeonPlayerCellChangeNotify, func() any { return new(proto.RogueDungeonPlayerCellChangeNotify) })
	c.regMsg(RogueFinishRepairReq, func() any { return new(proto.RogueFinishRepairReq) })
	c.regMsg(RogueFinishRepairRsp, func() any { return new(proto.RogueFinishRepairRsp) })
	c.regMsg(RogueHealAvatarsReq, func() any { return new(proto.RogueHealAvatarsReq) })
	c.regMsg(RogueHealAvatarsRsp, func() any { return new(proto.RogueHealAvatarsRsp) })
	c.regMsg(RogueResumeDungeonReq, func() any { return new(proto.RogueResumeDungeonReq) })
	c.regMsg(RogueResumeDungeonRsp, func() any { return new(proto.RogueResumeDungeonRsp) })
	c.regMsg(RogueSwitchAvatarReq, func() any { return new(proto.RogueSwitchAvatarReq) })
	c.regMsg(RogueSwitchAvatarRsp, func() any { return new(proto.RogueSwitchAvatarRsp) })
	c.regMsg(RoguelikeCardGachaNotify, func() any { return new(proto.RoguelikeCardGachaNotify) })
	c.regMsg(RoguelikeEffectDataNotify, func() any { return new(proto.RoguelikeEffectDataNotify) })
	c.regMsg(RoguelikeEffectViewReq, func() any { return new(proto.RoguelikeEffectViewReq) })
	c.regMsg(RoguelikeEffectViewRsp, func() any { return new(proto.RoguelikeEffectViewRsp) })
	c.regMsg(RoguelikeGiveUpReq, func() any { return new(proto.RoguelikeGiveUpReq) })
	c.regMsg(RoguelikeGiveUpRsp, func() any { return new(proto.RoguelikeGiveUpRsp) })
	c.regMsg(RoguelikeMistClearNotify, func() any { return new(proto.RoguelikeMistClearNotify) })
	c.regMsg(RoguelikeRefreshCardCostUpdateNotify, func() any { return new(proto.RoguelikeRefreshCardCostUpdateNotify) })
	c.regMsg(RoguelikeResourceBonusPropUpdateNotify, func() any { return new(proto.RoguelikeResourceBonusPropUpdateNotify) })
	c.regMsg(RoguelikeRuneRecordUpdateNotify, func() any { return new(proto.RoguelikeRuneRecordUpdateNotify) })
	c.regMsg(RoguelikeSelectAvatarAndEnterDungeonReq, func() any { return new(proto.RoguelikeSelectAvatarAndEnterDungeonReq) })
	c.regMsg(RoguelikeSelectAvatarAndEnterDungeonRsp, func() any { return new(proto.RoguelikeSelectAvatarAndEnterDungeonRsp) })
	c.regMsg(RoguelikeTakeStageFirstPassRewardReq, func() any { return new(proto.RoguelikeTakeStageFirstPassRewardReq) })
	c.regMsg(RoguelikeTakeStageFirstPassRewardRsp, func() any { return new(proto.RoguelikeTakeStageFirstPassRewardRsp) })
	c.regMsg(SalesmanDeliverItemReq, func() any { return new(proto.SalesmanDeliverItemReq) })
	c.regMsg(SalesmanDeliverItemRsp, func() any { return new(proto.SalesmanDeliverItemRsp) })
	c.regMsg(SalesmanTakeRewardReq, func() any { return new(proto.SalesmanTakeRewardReq) })
	c.regMsg(SalesmanTakeRewardRsp, func() any { return new(proto.SalesmanTakeRewardRsp) })
	c.regMsg(SalesmanTakeSpecialRewardReq, func() any { return new(proto.SalesmanTakeSpecialRewardReq) })
	c.regMsg(SalesmanTakeSpecialRewardRsp, func() any { return new(proto.SalesmanTakeSpecialRewardRsp) })
	c.regMsg(SalvageEscortRestartReq, func() any { return new(proto.SalvageEscortRestartReq) })
	c.regMsg(SalvageEscortRestartRsp, func() any { return new(proto.SalvageEscortRestartRsp) })
	c.regMsg(SalvageEscortSettleNotify, func() any { return new(proto.SalvageEscortSettleNotify) })
	c.regMsg(SalvagePreventRestartReq, func() any { return new(proto.SalvagePreventRestartReq) })
	c.regMsg(SalvagePreventRestartRsp, func() any { return new(proto.SalvagePreventRestartRsp) })
	c.regMsg(SalvagePreventSettleNotify, func() any { return new(proto.SalvagePreventSettleNotify) })
	c.regMsg(SaveCoopDialogReq, func() any { return new(proto.SaveCoopDialogReq) })
	c.regMsg(SaveCoopDialogRsp, func() any { return new(proto.SaveCoopDialogRsp) })
	c.regMsg(SaveCustomDungeonRoomReq, func() any { return new(proto.SaveCustomDungeonRoomReq) })
	c.regMsg(SaveCustomDungeonRoomRsp, func() any { return new(proto.SaveCustomDungeonRoomRsp) })
	c.regMsg(SaveMainCoopReq, func() any { return new(proto.SaveMainCoopReq) })
	c.regMsg(SaveMainCoopRsp, func() any { return new(proto.SaveMainCoopRsp) })
	c.regMsg(SaveUgcReq, func() any { return new(proto.SaveUgcReq) })
	c.regMsg(SaveUgcRsp, func() any { return new(proto.SaveUgcRsp) })
	c.regMsg(SceneAreaUnlockNotify, func() any { return new(proto.SceneAreaUnlockNotify) })
	c.regMsg(SceneAreaWeatherNotify, func() any { return new(proto.SceneAreaWeatherNotify) })
	c.regMsg(SceneAudioNotify, func() any { return new(proto.SceneAudioNotify) })
	c.regMsg(SceneAvatarStaminaStepReq, func() any { return new(proto.SceneAvatarStaminaStepReq) })
	c.regMsg(SceneAvatarStaminaStepRsp, func() any { return new(proto.SceneAvatarStaminaStepRsp) })
	c.regMsg(SceneCreateEntityReq, func() any { return new(proto.SceneCreateEntityReq) })
	c.regMsg(SceneCreateEntityRsp, func() any { return new(proto.SceneCreateEntityRsp) })
	c.regMsg(SceneDataNotify, func() any { return new(proto.SceneDataNotify) })
	c.regMsg(SceneDestroyEntityReq, func() any { return new(proto.SceneDestroyEntityReq) })
	c.regMsg(SceneDestroyEntityRsp, func() any { return new(proto.SceneDestroyEntityRsp) })
	c.regMsg(SceneEntitiesMoveCombineNotify, func() any { return new(proto.SceneEntitiesMoveCombineNotify) })
	c.regMsg(SceneEntitiesMovesReq, func() any { return new(proto.SceneEntitiesMovesReq) })
	c.regMsg(SceneEntitiesMovesRsp, func() any { return new(proto.SceneEntitiesMovesRsp) })
	c.regMsg(SceneEntityAppearNotify, func() any { return new(proto.SceneEntityAppearNotify) })
	c.regMsg(SceneEntityDisappearNotify, func() any { return new(proto.SceneEntityDisappearNotify) })
	c.regMsg(SceneEntityDrownReq, func() any { return new(proto.SceneEntityDrownReq) })
	c.regMsg(SceneEntityDrownRsp, func() any { return new(proto.SceneEntityDrownRsp) })
	c.regMsg(SceneEntityMoveNotify, func() any { return new(proto.SceneEntityMoveNotify) })
	c.regMsg(SceneEntityMoveReq, func() any { return new(proto.SceneEntityMoveReq) })
	c.regMsg(SceneEntityMoveRsp, func() any { return new(proto.SceneEntityMoveRsp) })
	c.regMsg(SceneEntityUpdateNotify, func() any { return new(proto.SceneEntityUpdateNotify) })
	c.regMsg(SceneForceLockNotify, func() any { return new(proto.SceneForceLockNotify) })
	c.regMsg(SceneForceUnlockNotify, func() any { return new(proto.SceneForceUnlockNotify) })
	c.regMsg(SceneGalleryInfoNotify, func() any { return new(proto.SceneGalleryInfoNotify) })
	c.regMsg(SceneGalleryVintageHuntingSettleNotify, func() any { return new(proto.SceneGalleryVintageHuntingSettleNotify) })
	c.regMsg(SceneInitFinishReq, func() any { return new(proto.SceneInitFinishReq) })
	c.regMsg(SceneInitFinishRsp, func() any { return new(proto.SceneInitFinishRsp) })
	c.regMsg(SceneKickPlayerNotify, func() any { return new(proto.SceneKickPlayerNotify) })
	c.regMsg(SceneKickPlayerReq, func() any { return new(proto.SceneKickPlayerReq) })
	c.regMsg(SceneKickPlayerRsp, func() any { return new(proto.SceneKickPlayerRsp) })
	c.regMsg(ScenePlayBattleInfoListNotify, func() any { return new(proto.ScenePlayBattleInfoListNotify) })
	c.regMsg(ScenePlayBattleInfoNotify, func() any { return new(proto.ScenePlayBattleInfoNotify) })
	c.regMsg(ScenePlayBattleInterruptNotify, func() any { return new(proto.ScenePlayBattleInterruptNotify) })
	c.regMsg(ScenePlayBattleResultNotify, func() any { return new(proto.ScenePlayBattleResultNotify) })
	c.regMsg(ScenePlayBattleUidOpNotify, func() any { return new(proto.ScenePlayBattleUidOpNotify) })
	c.regMsg(ScenePlayGuestReplyInviteReq, func() any { return new(proto.ScenePlayGuestReplyInviteReq) })
	c.regMsg(ScenePlayGuestReplyInviteRsp, func() any { return new(proto.ScenePlayGuestReplyInviteRsp) })
	c.regMsg(ScenePlayGuestReplyNotify, func() any { return new(proto.ScenePlayGuestReplyNotify) })
	c.regMsg(ScenePlayInfoListNotify, func() any { return new(proto.ScenePlayInfoListNotify) })
	c.regMsg(ScenePlayInviteResultNotify, func() any { return new(proto.ScenePlayInviteResultNotify) })
	c.regMsg(ScenePlayOutofRegionNotify, func() any { return new(proto.ScenePlayOutofRegionNotify) })
	c.regMsg(ScenePlayOwnerCheckReq, func() any { return new(proto.ScenePlayOwnerCheckReq) })
	c.regMsg(ScenePlayOwnerCheckRsp, func() any { return new(proto.ScenePlayOwnerCheckRsp) })
	c.regMsg(ScenePlayOwnerInviteNotify, func() any { return new(proto.ScenePlayOwnerInviteNotify) })
	c.regMsg(ScenePlayOwnerStartInviteReq, func() any { return new(proto.ScenePlayOwnerStartInviteReq) })
	c.regMsg(ScenePlayOwnerStartInviteRsp, func() any { return new(proto.ScenePlayOwnerStartInviteRsp) })
	c.regMsg(ScenePlayerBackgroundAvatarRefreshNotify, func() any { return new(proto.ScenePlayerBackgroundAvatarRefreshNotify) })
	c.regMsg(ScenePlayerInfoNotify, func() any { return new(proto.ScenePlayerInfoNotify) })
	c.regMsg(ScenePlayerLocationNotify, func() any { return new(proto.ScenePlayerLocationNotify) })
	c.regMsg(ScenePlayerSoundNotify, func() any { return new(proto.ScenePlayerSoundNotify) })
	c.regMsg(ScenePointUnlockNotify, func() any { return new(proto.ScenePointUnlockNotify) })
	c.regMsg(SceneRouteChangeNotify, func() any { return new(proto.SceneRouteChangeNotify) })
	c.regMsg(SceneTeamUpdateNotify, func() any { return new(proto.SceneTeamUpdateNotify) })
	c.regMsg(SceneTimeNotify, func() any { return new(proto.SceneTimeNotify) })
	c.regMsg(SceneTransToPointReq, func() any { return new(proto.SceneTransToPointReq) })
	c.regMsg(SceneTransToPointRsp, func() any { return new(proto.SceneTransToPointRsp) })
	c.regMsg(SeaLampCoinNotify, func() any { return new(proto.SeaLampCoinNotify) })
	c.regMsg(SeaLampContributeItemReq, func() any { return new(proto.SeaLampContributeItemReq) })
	c.regMsg(SeaLampContributeItemRsp, func() any { return new(proto.SeaLampContributeItemRsp) })
	c.regMsg(SeaLampFlyLampNotify, func() any { return new(proto.SeaLampFlyLampNotify) })
	c.regMsg(SeaLampFlyLampReq, func() any { return new(proto.SeaLampFlyLampReq) })
	c.regMsg(SeaLampFlyLampRsp, func() any { return new(proto.SeaLampFlyLampRsp) })
	c.regMsg(SeaLampPopularityNotify, func() any { return new(proto.SeaLampPopularityNotify) })
	c.regMsg(SeaLampTakeContributionRewardReq, func() any { return new(proto.SeaLampTakeContributionRewardReq) })
	c.regMsg(SeaLampTakeContributionRewardRsp, func() any { return new(proto.SeaLampTakeContributionRewardRsp) })
	c.regMsg(SeaLampTakePhaseRewardReq, func() any { return new(proto.SeaLampTakePhaseRewardReq) })
	c.regMsg(SeaLampTakePhaseRewardRsp, func() any { return new(proto.SeaLampTakePhaseRewardRsp) })
	c.regMsg(SealBattleBeginNotify, func() any { return new(proto.SealBattleBeginNotify) })
	c.regMsg(SealBattleEndNotify, func() any { return new(proto.SealBattleEndNotify) })
	c.regMsg(SealBattleProgressNotify, func() any { return new(proto.SealBattleProgressNotify) })
	c.regMsg(SearchCustomDungeonReq, func() any { return new(proto.SearchCustomDungeonReq) })
	c.regMsg(SearchCustomDungeonRsp, func() any { return new(proto.SearchCustomDungeonRsp) })
	c.regMsg(SeeMonsterReq, func() any { return new(proto.SeeMonsterReq) })
	c.regMsg(SeeMonsterRsp, func() any { return new(proto.SeeMonsterRsp) })
	c.regMsg(SelectAsterMidDifficultyReq, func() any { return new(proto.SelectAsterMidDifficultyReq) })
	c.regMsg(SelectAsterMidDifficultyRsp, func() any { return new(proto.SelectAsterMidDifficultyRsp) })
	c.regMsg(SelectEffigyChallengeConditionReq, func() any { return new(proto.SelectEffigyChallengeConditionReq) })
	c.regMsg(SelectEffigyChallengeConditionRsp, func() any { return new(proto.SelectEffigyChallengeConditionRsp) })
	c.regMsg(SelectRoguelikeDungeonCardReq, func() any { return new(proto.SelectRoguelikeDungeonCardReq) })
	c.regMsg(SelectRoguelikeDungeonCardRsp, func() any { return new(proto.SelectRoguelikeDungeonCardRsp) })
	c.regMsg(SelectWorktopOptionReq, func() any { return new(proto.SelectWorktopOptionReq) })
	c.regMsg(SelectWorktopOptionRsp, func() any { return new(proto.SelectWorktopOptionRsp) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(ServerAnnounceRevokeNotify, func() any { return new(proto.ServerAnnounceRevokeNotify) })
	c.regMsg(ServerBuffChangeNotify, func() any { return new(proto.ServerBuffChangeNotify) })
	c.regMsg(ServerCombatEndNotify, func() any { return new(proto.ServerCombatEndNotify) })
	c.regMsg(ServerCondMeetQuestListUpdateNotify, func() any { return new(proto.ServerCondMeetQuestListUpdateNotify) })
	c.regMsg(ServerDisconnectClientNotify, func() any { return new(proto.ServerDisconnectClientNotify) })
	c.regMsg(ServerGlobalValueChangeNotify, func() any { return new(proto.ServerGlobalValueChangeNotify) })
	c.regMsg(ServerLogNotify, func() any { return new(proto.ServerLogNotify) })
	c.regMsg(ServerMessageNotify, func() any { return new(proto.ServerMessageNotify) })
	c.regMsg(ServerTimeNotify, func() any { return new(proto.ServerTimeNotify) })
	c.regMsg(ServerTryCancelGeneralMatchNotify, func() any { return new(proto.ServerTryCancelGeneralMatchNotify) })
	c.regMsg(ServerUpdateGlobalValueNotify, func() any { return new(proto.ServerUpdateGlobalValueNotify) })
	c.regMsg(SetBattlePassViewedReq, func() any { return new(proto.SetBattlePassViewedReq) })
	c.regMsg(SetBattlePassViewedRsp, func() any { return new(proto.SetBattlePassViewedRsp) })
	c.regMsg(SetChatEmojiCollectionReq, func() any { return new(proto.SetChatEmojiCollectionReq) })
	c.regMsg(SetChatEmojiCollectionRsp, func() any { return new(proto.SetChatEmojiCollectionRsp) })
	c.regMsg(SetCodexPushtipsReadReq, func() any { return new(proto.SetCodexPushtipsReadReq) })
	c.regMsg(SetCodexPushtipsReadRsp, func() any { return new(proto.SetCodexPushtipsReadRsp) })
	c.regMsg(SetCoopChapterViewedReq, func() any { return new(proto.SetCoopChapterViewedReq) })
	c.regMsg(SetCoopChapterViewedRsp, func() any { return new(proto.SetCoopChapterViewedRsp) })
	c.regMsg(SetCurExpeditionChallengeIdReq, func() any { return new(proto.SetCurExpeditionChallengeIdReq) })
	c.regMsg(SetCurExpeditionChallengeIdRsp, func() any { return new(proto.SetCurExpeditionChallengeIdRsp) })
	c.regMsg(SetEntityClientDataNotify, func() any { return new(proto.SetEntityClientDataNotify) })
	c.regMsg(SetEquipLockStateReq, func() any { return new(proto.SetEquipLockStateReq) })
	c.regMsg(SetEquipLockStateRsp, func() any { return new(proto.SetEquipLockStateRsp) })
	c.regMsg(SetFriendEnterHomeOptionReq, func() any { return new(proto.SetFriendEnterHomeOptionReq) })
	c.regMsg(SetFriendEnterHomeOptionRsp, func() any { return new(proto.SetFriendEnterHomeOptionRsp) })
	c.regMsg(SetFriendRemarkNameReq, func() any { return new(proto.SetFriendRemarkNameReq) })
	c.regMsg(SetFriendRemarkNameRsp, func() any { return new(proto.SetFriendRemarkNameRsp) })
	c.regMsg(SetH5ActivityRedDotTimestampReq, func() any { return new(proto.SetH5ActivityRedDotTimestampReq) })
	c.regMsg(SetH5ActivityRedDotTimestampRsp, func() any { return new(proto.SetH5ActivityRedDotTimestampRsp) })
	c.regMsg(SetIsAutoUnlockSpecificEquipReq, func() any { return new(proto.SetIsAutoUnlockSpecificEquipReq) })
	c.regMsg(SetIsAutoUnlockSpecificEquipRsp, func() any { return new(proto.SetIsAutoUnlockSpecificEquipRsp) })
	c.regMsg(SetLimitOptimizationNotify, func() any { return new(proto.SetLimitOptimizationNotify) })
	c.regMsg(SetNameCardReq, func() any { return new(proto.SetNameCardReq) })
	c.regMsg(SetNameCardRsp, func() any { return new(proto.SetNameCardRsp) })
	c.regMsg(SetOpenStateReq, func() any { return new(proto.SetOpenStateReq) })
	c.regMsg(SetOpenStateRsp, func() any { return new(proto.SetOpenStateRsp) })
	c.regMsg(SetPlayerBirthdayReq, func() any { return new(proto.SetPlayerBirthdayReq) })
	c.regMsg(SetPlayerBirthdayRsp, func() any { return new(proto.SetPlayerBirthdayRsp) })
	c.regMsg(SetPlayerBornDataReq, func() any { return new(proto.SetPlayerBornDataReq) })
	c.regMsg(SetPlayerBornDataRsp, func() any { return new(proto.SetPlayerBornDataRsp) })
	c.regMsg(SetPlayerHeadImageReq, func() any { return new(proto.SetPlayerHeadImageReq) })
	c.regMsg(SetPlayerHeadImageRsp, func() any { return new(proto.SetPlayerHeadImageRsp) })
	c.regMsg(SetPlayerNameReq, func() any { return new(proto.SetPlayerNameReq) })
	c.regMsg(SetPlayerNameRsp, func() any { return new(proto.SetPlayerNameRsp) })
	c.regMsg(SetPlayerPropReq, func() any { return new(proto.SetPlayerPropReq) })
	c.regMsg(SetPlayerPropRsp, func() any { return new(proto.SetPlayerPropRsp) })
	c.regMsg(SetPlayerSignatureReq, func() any { return new(proto.SetPlayerSignatureReq) })
	c.regMsg(SetPlayerSignatureRsp, func() any { return new(proto.SetPlayerSignatureRsp) })
	c.regMsg(SetSceneWeatherAreaReq, func() any { return new(proto.SetSceneWeatherAreaReq) })
	c.regMsg(SetSceneWeatherAreaRsp, func() any { return new(proto.SetSceneWeatherAreaRsp) })
	c.regMsg(SetUpAvatarTeamReq, func() any { return new(proto.SetUpAvatarTeamReq) })
	c.regMsg(SetUpAvatarTeamRsp, func() any { return new(proto.SetUpAvatarTeamRsp) })
	c.regMsg(SetUpLunchBoxWidgetReq, func() any { return new(proto.SetUpLunchBoxWidgetReq) })
	c.regMsg(SetUpLunchBoxWidgetRsp, func() any { return new(proto.SetUpLunchBoxWidgetRsp) })
	c.regMsg(SetWidgetSlotReq, func() any { return new(proto.SetWidgetSlotReq) })
	c.regMsg(SetWidgetSlotRsp, func() any { return new(proto.SetWidgetSlotRsp) })
	c.regMsg(ShowClientGuideNotify, func() any { return new(proto.ShowClientGuideNotify) })
	c.regMsg(ShowClientTutorialNotify, func() any { return new(proto.ShowClientTutorialNotify) })
	c.regMsg(ShowCommonTipsNotify, func() any { return new(proto.ShowCommonTipsNotify) })
	c.regMsg(ShowMessageNotify, func() any { return new(proto.ShowMessageNotify) })
	c.regMsg(ShowTemplateReminderNotify, func() any { return new(proto.ShowTemplateReminderNotify) })
	c.regMsg(SignInInfoReq, func() any { return new(proto.SignInInfoReq) })
	c.regMsg(SignInInfoRsp, func() any { return new(proto.SignInInfoRsp) })
	c.regMsg(SignatureAuditConfigNotify, func() any { return new(proto.SignatureAuditConfigNotify) })
	c.regMsg(SkyCrystalDetectorDataUpdateNotify, func() any { return new(proto.SkyCrystalDetectorDataUpdateNotify) })
	c.regMsg(SocialDataNotify, func() any { return new(proto.SocialDataNotify) })
	c.regMsg(SpiceActivityFinishMakeSpiceReq, func() any { return new(proto.SpiceActivityFinishMakeSpiceReq) })
	c.regMsg(SpiceActivityFinishMakeSpiceRsp, func() any { return new(proto.SpiceActivityFinishMakeSpiceRsp) })
	c.regMsg(SpiceActivityGivingRecordNotify, func() any { return new(proto.SpiceActivityGivingRecordNotify) })
	c.regMsg(SpiceActivityProcessFoodReq, func() any { return new(proto.SpiceActivityProcessFoodReq) })
	c.regMsg(SpiceActivityProcessFoodRsp, func() any { return new(proto.SpiceActivityProcessFoodRsp) })
	c.regMsg(SpringUseReq, func() any { return new(proto.SpringUseReq) })
	c.regMsg(SpringUseRsp, func() any { return new(proto.SpringUseRsp) })
	c.regMsg(StartArenaChallengeLevelReq, func() any { return new(proto.StartArenaChallengeLevelReq) })
	c.regMsg(StartArenaChallengeLevelRsp, func() any { return new(proto.StartArenaChallengeLevelRsp) })
	c.regMsg(StartBuoyantCombatGalleryReq, func() any { return new(proto.StartBuoyantCombatGalleryReq) })
	c.regMsg(StartBuoyantCombatGalleryRsp, func() any { return new(proto.StartBuoyantCombatGalleryRsp) })
	c.regMsg(StartCoopPointReq, func() any { return new(proto.StartCoopPointReq) })
	c.regMsg(StartCoopPointRsp, func() any { return new(proto.StartCoopPointRsp) })
	c.regMsg(StartEffigyChallengeReq, func() any { return new(proto.StartEffigyChallengeReq) })
	c.regMsg(StartEffigyChallengeRsp, func() any { return new(proto.StartEffigyChallengeRsp) })
	c.regMsg(StartFishingReq, func() any { return new(proto.StartFishingReq) })
	c.regMsg(StartFishingRsp, func() any { return new(proto.StartFishingRsp) })
	c.regMsg(StartRogueDiaryPlayReq, func() any { return new(proto.StartRogueDiaryPlayReq) })
	c.regMsg(StartRogueDiaryPlayRsp, func() any { return new(proto.StartRogueDiaryPlayRsp) })
	c.regMsg(StartRogueDiaryRoomReq, func() any { return new(proto.StartRogueDiaryRoomReq) })
	c.regMsg(StartRogueDiaryRoomRsp, func() any { return new(proto.StartRogueDiaryRoomRsp) })
	c.regMsg(StartRogueEliteCellChallengeReq, func() any { return new(proto.StartRogueEliteCellChallengeReq) })
	c.regMsg(StartRogueEliteCellChallengeRsp, func() any { return new(proto.StartRogueEliteCellChallengeRsp) })
	c.regMsg(StartRogueNormalCellChallengeReq, func() any { return new(proto.StartRogueNormalCellChallengeReq) })
	c.regMsg(StartRogueNormalCellChallengeRsp, func() any { return new(proto.StartRogueNormalCellChallengeRsp) })
	c.regMsg(StopReminderNotify, func() any { return new(proto.StopReminderNotify) })
	c.regMsg(StoreCustomDungeonReq, func() any { return new(proto.StoreCustomDungeonReq) })
	c.regMsg(StoreCustomDungeonRsp, func() any { return new(proto.StoreCustomDungeonRsp) })
	c.regMsg(StoreItemChangeNotify, func() any { return new(proto.StoreItemChangeNotify) })
	c.regMsg(StoreItemDelNotify, func() any { return new(proto.StoreItemDelNotify) })
	c.regMsg(StoreWeightLimitNotify, func() any { return new(proto.StoreWeightLimitNotify) })
	c.regMsg(SubmitInferenceWordReq, func() any { return new(proto.SubmitInferenceWordReq) })
	c.regMsg(SubmitInferenceWordRsp, func() any { return new(proto.SubmitInferenceWordRsp) })
	c.regMsg(SummerTimeFloatSignalPositionNotify, func() any { return new(proto.SummerTimeFloatSignalPositionNotify) })
	c.regMsg(SummerTimeFloatSignalUpdateNotify, func() any { return new(proto.SummerTimeFloatSignalUpdateNotify) })
	c.regMsg(SummerTimeSprintBoatRestartReq, func() any { return new(proto.SummerTimeSprintBoatRestartReq) })
	c.regMsg(SummerTimeSprintBoatRestartRsp, func() any { return new(proto.SummerTimeSprintBoatRestartRsp) })
	c.regMsg(SummerTimeSprintBoatSettleNotify, func() any { return new(proto.SummerTimeSprintBoatSettleNotify) })
	c.regMsg(SummerTimeV2BoatSettleNotify, func() any { return new(proto.SummerTimeV2BoatSettleNotify) })
	c.regMsg(SummerTimeV2RestartBoatGalleryReq, func() any { return new(proto.SummerTimeV2RestartBoatGalleryReq) })
	c.regMsg(SummerTimeV2RestartBoatGalleryRsp, func() any { return new(proto.SummerTimeV2RestartBoatGalleryRsp) })
	c.regMsg(SummerTimeV2RestartDungeonReq, func() any { return new(proto.SummerTimeV2RestartDungeonReq) })
	c.regMsg(SummerTimeV2RestartDungeonRsp, func() any { return new(proto.SummerTimeV2RestartDungeonRsp) })
	c.regMsg(SumoDungeonSettleNotify, func() any { return new(proto.SumoDungeonSettleNotify) })
	c.regMsg(SumoEnterDungeonNotify, func() any { return new(proto.SumoEnterDungeonNotify) })
	c.regMsg(SumoLeaveDungeonNotify, func() any { return new(proto.SumoLeaveDungeonNotify) })
	c.regMsg(SumoRestartDungeonReq, func() any { return new(proto.SumoRestartDungeonReq) })
	c.regMsg(SumoRestartDungeonRsp, func() any { return new(proto.SumoRestartDungeonRsp) })
	c.regMsg(SumoSaveTeamReq, func() any { return new(proto.SumoSaveTeamReq) })
	c.regMsg(SumoSaveTeamRsp, func() any { return new(proto.SumoSaveTeamRsp) })
	c.regMsg(SumoSelectTeamAndEnterDungeonReq, func() any { return new(proto.SumoSelectTeamAndEnterDungeonReq) })
	c.regMsg(SumoSelectTeamAndEnterDungeonRsp, func() any { return new(proto.SumoSelectTeamAndEnterDungeonRsp) })
	c.regMsg(SumoSetNoSwitchPunishTimeNotify, func() any { return new(proto.SumoSetNoSwitchPunishTimeNotify) })
	c.regMsg(SumoSwitchTeamReq, func() any { return new(proto.SumoSwitchTeamReq) })
	c.regMsg(SumoSwitchTeamRsp, func() any { return new(proto.SumoSwitchTeamRsp) })
	c.regMsg(SyncScenePlayTeamEntityNotify, func() any { return new(proto.SyncScenePlayTeamEntityNotify) })
	c.regMsg(SyncTeamEntityNotify, func() any { return new(proto.SyncTeamEntityNotify) })
	c.regMsg(TakeAchievementGoalRewardReq, func() any { return new(proto.TakeAchievementGoalRewardReq) })
	c.regMsg(TakeAchievementGoalRewardRsp, func() any { return new(proto.TakeAchievementGoalRewardRsp) })
	c.regMsg(TakeAchievementRewardReq, func() any { return new(proto.TakeAchievementRewardReq) })
	c.regMsg(TakeAchievementRewardRsp, func() any { return new(proto.TakeAchievementRewardRsp) })
	c.regMsg(TakeAsterSpecialRewardReq, func() any { return new(proto.TakeAsterSpecialRewardReq) })
	c.regMsg(TakeAsterSpecialRewardRsp, func() any { return new(proto.TakeAsterSpecialRewardRsp) })
	c.regMsg(TakeBackGivingItemReq, func() any { return new(proto.TakeBackGivingItemReq) })
	c.regMsg(TakeBackGivingItemRsp, func() any { return new(proto.TakeBackGivingItemRsp) })
	c.regMsg(TakeBattlePassMissionPointReq, func() any { return new(proto.TakeBattlePassMissionPointReq) })
	c.regMsg(TakeBattlePassMissionPointRsp, func() any { return new(proto.TakeBattlePassMissionPointRsp) })
	c.regMsg(TakeBattlePassRewardReq, func() any { return new(proto.TakeBattlePassRewardReq) })
	c.regMsg(TakeBattlePassRewardRsp, func() any { return new(proto.TakeBattlePassRewardRsp) })
	c.regMsg(TakeCityReputationExploreRewardReq, func() any { return new(proto.TakeCityReputationExploreRewardReq) })
	c.regMsg(TakeCityReputationExploreRewardRsp, func() any { return new(proto.TakeCityReputationExploreRewardRsp) })
	c.regMsg(TakeCityReputationLevelRewardReq, func() any { return new(proto.TakeCityReputationLevelRewardReq) })
	c.regMsg(TakeCityReputationLevelRewardRsp, func() any { return new(proto.TakeCityReputationLevelRewardRsp) })
	c.regMsg(TakeCityReputationParentQuestReq, func() any { return new(proto.TakeCityReputationParentQuestReq) })
	c.regMsg(TakeCityReputationParentQuestRsp, func() any { return new(proto.TakeCityReputationParentQuestRsp) })
	c.regMsg(TakeCompoundOutputReq, func() any { return new(proto.TakeCompoundOutputReq) })
	c.regMsg(TakeCompoundOutputRsp, func() any { return new(proto.TakeCompoundOutputRsp) })
	c.regMsg(TakeCoopRewardReq, func() any { return new(proto.TakeCoopRewardReq) })
	c.regMsg(TakeCoopRewardRsp, func() any { return new(proto.TakeCoopRewardRsp) })
	c.regMsg(TakeDeliveryDailyRewardReq, func() any { return new(proto.TakeDeliveryDailyRewardReq) })
	c.regMsg(TakeDeliveryDailyRewardRsp, func() any { return new(proto.TakeDeliveryDailyRewardRsp) })
	c.regMsg(TakeEffigyFirstPassRewardReq, func() any { return new(proto.TakeEffigyFirstPassRewardReq) })
	c.regMsg(TakeEffigyFirstPassRewardRsp, func() any { return new(proto.TakeEffigyFirstPassRewardRsp) })
	c.regMsg(TakeEffigyRewardReq, func() any { return new(proto.TakeEffigyRewardReq) })
	c.regMsg(TakeEffigyRewardRsp, func() any { return new(proto.TakeEffigyRewardRsp) })
	c.regMsg(TakeFirstShareRewardReq, func() any { return new(proto.TakeFirstShareRewardReq) })
	c.regMsg(TakeFirstShareRewardRsp, func() any { return new(proto.TakeFirstShareRewardRsp) })
	c.regMsg(TakeFurnitureMakeReq, func() any { return new(proto.TakeFurnitureMakeReq) })
	c.regMsg(TakeFurnitureMakeRsp, func() any { return new(proto.TakeFurnitureMakeRsp) })
	c.regMsg(TakeHuntingOfferReq, func() any { return new(proto.TakeHuntingOfferReq) })
	c.regMsg(TakeHuntingOfferRsp, func() any { return new(proto.TakeHuntingOfferRsp) })
	c.regMsg(TakeInvestigationRewardReq, func() any { return new(proto.TakeInvestigationRewardReq) })
	c.regMsg(TakeInvestigationRewardRsp, func() any { return new(proto.TakeInvestigationRewardRsp) })
	c.regMsg(TakeInvestigationTargetRewardReq, func() any { return new(proto.TakeInvestigationTargetRewardReq) })
	c.regMsg(TakeInvestigationTargetRewardRsp, func() any { return new(proto.TakeInvestigationTargetRewardRsp) })
	c.regMsg(TakeMaterialDeleteReturnReq, func() any { return new(proto.TakeMaterialDeleteReturnReq) })
	c.regMsg(TakeMaterialDeleteReturnRsp, func() any { return new(proto.TakeMaterialDeleteReturnRsp) })
	c.regMsg(TakeOfferingLevelRewardReq, func() any { return new(proto.TakeOfferingLevelRewardReq) })
	c.regMsg(TakeOfferingLevelRewardRsp, func() any { return new(proto.TakeOfferingLevelRewardRsp) })
	c.regMsg(TakePlayerLevelRewardReq, func() any { return new(proto.TakePlayerLevelRewardReq) })
	c.regMsg(TakePlayerLevelRewardRsp, func() any { return new(proto.TakePlayerLevelRewardRsp) })
	c.regMsg(TakeRegionSearchRewardReq, func() any { return new(proto.TakeRegionSearchRewardReq) })
	c.regMsg(TakeRegionSearchRewardRsp, func() any { return new(proto.TakeRegionSearchRewardRsp) })
	c.regMsg(TakeResinCardDailyRewardReq, func() any { return new(proto.TakeResinCardDailyRewardReq) })
	c.regMsg(TakeResinCardDailyRewardRsp, func() any { return new(proto.TakeResinCardDailyRewardRsp) })
	c.regMsg(TakeReunionFirstGiftRewardReq, func() any { return new(proto.TakeReunionFirstGiftRewardReq) })
	c.regMsg(TakeReunionFirstGiftRewardRsp, func() any { return new(proto.TakeReunionFirstGiftRewardRsp) })
	c.regMsg(TakeReunionMissionRewardReq, func() any { return new(proto.TakeReunionMissionRewardReq) })
	c.regMsg(TakeReunionMissionRewardRsp, func() any { return new(proto.TakeReunionMissionRewardRsp) })
	c.regMsg(TakeReunionSignInRewardReq, func() any { return new(proto.TakeReunionSignInRewardReq) })
	c.regMsg(TakeReunionSignInRewardRsp, func() any { return new(proto.TakeReunionSignInRewardRsp) })
	c.regMsg(TakeReunionWatcherRewardReq, func() any { return new(proto.TakeReunionWatcherRewardReq) })
	c.regMsg(TakeReunionWatcherRewardRsp, func() any { return new(proto.TakeReunionWatcherRewardRsp) })
	c.regMsg(TakeoffEquipReq, func() any { return new(proto.TakeoffEquipReq) })
	c.regMsg(TakeoffEquipRsp, func() any { return new(proto.TakeoffEquipRsp) })
	c.regMsg(TanukiTravelFinishGuideQuestNotify, func() any { return new(proto.TanukiTravelFinishGuideQuestNotify) })
	c.regMsg(TaskVarNotify, func() any { return new(proto.TaskVarNotify) })
	c.regMsg(TeamResonanceChangeNotify, func() any { return new(proto.TeamResonanceChangeNotify) })
	c.regMsg(ToTheMoonAddObstacleReq, func() any { return new(proto.ToTheMoonAddObstacleReq) })
	c.regMsg(ToTheMoonAddObstacleRsp, func() any { return new(proto.ToTheMoonAddObstacleRsp) })
	c.regMsg(ToTheMoonEnterSceneReq, func() any { return new(proto.ToTheMoonEnterSceneReq) })
	c.regMsg(ToTheMoonEnterSceneRsp, func() any { return new(proto.ToTheMoonEnterSceneRsp) })
	c.regMsg(ToTheMoonObstaclesModifyNotify, func() any { return new(proto.ToTheMoonObstaclesModifyNotify) })
	c.regMsg(ToTheMoonPingNotify, func() any { return new(proto.ToTheMoonPingNotify) })
	c.regMsg(ToTheMoonQueryPathReq, func() any { return new(proto.ToTheMoonQueryPathReq) })
	c.regMsg(ToTheMoonQueryPathRsp, func() any { return new(proto.ToTheMoonQueryPathRsp) })
	c.regMsg(ToTheMoonRemoveObstacleReq, func() any { return new(proto.ToTheMoonRemoveObstacleReq) })
	c.regMsg(ToTheMoonRemoveObstacleRsp, func() any { return new(proto.ToTheMoonRemoveObstacleRsp) })
	c.regMsg(TowerAllDataReq, func() any { return new(proto.TowerAllDataReq) })
	c.regMsg(TowerAllDataRsp, func() any { return new(proto.TowerAllDataRsp) })
	c.regMsg(TowerBriefDataNotify, func() any { return new(proto.TowerBriefDataNotify) })
	c.regMsg(TowerBuffSelectReq, func() any { return new(proto.TowerBuffSelectReq) })
	c.regMsg(TowerBuffSelectRsp, func() any { return new(proto.TowerBuffSelectRsp) })
	c.regMsg(TowerCurLevelRecordChangeNotify, func() any { return new(proto.TowerCurLevelRecordChangeNotify) })
	c.regMsg(TowerDailyRewardProgressChangeNotify, func() any { return new(proto.TowerDailyRewardProgressChangeNotify) })
	c.regMsg(TowerEnterLevelReq, func() any { return new(proto.TowerEnterLevelReq) })
	c.regMsg(TowerEnterLevelRsp, func() any { return new(proto.TowerEnterLevelRsp) })
	c.regMsg(TowerFloorRecordChangeNotify, func() any { return new(proto.TowerFloorRecordChangeNotify) })
	c.regMsg(TowerGetFloorStarRewardReq, func() any { return new(proto.TowerGetFloorStarRewardReq) })
	c.regMsg(TowerGetFloorStarRewardRsp, func() any { return new(proto.TowerGetFloorStarRewardRsp) })
	c.regMsg(TowerLevelEndNotify, func() any { return new(proto.TowerLevelEndNotify) })
	c.regMsg(TowerLevelStarCondNotify, func() any { return new(proto.TowerLevelStarCondNotify) })
	c.regMsg(TowerMiddleLevelChangeTeamNotify, func() any { return new(proto.TowerMiddleLevelChangeTeamNotify) })
	c.regMsg(TowerRecordHandbookReq, func() any { return new(proto.TowerRecordHandbookReq) })
	c.regMsg(TowerRecordHandbookRsp, func() any { return new(proto.TowerRecordHandbookRsp) })
	c.regMsg(TowerSurrenderReq, func() any { return new(proto.TowerSurrenderReq) })
	c.regMsg(TowerSurrenderRsp, func() any { return new(proto.TowerSurrenderRsp) })
	c.regMsg(TowerTeamSelectReq, func() any { return new(proto.TowerTeamSelectReq) })
	c.regMsg(TowerTeamSelectRsp, func() any { return new(proto.TowerTeamSelectRsp) })
	c.regMsg(TreasureMapBonusChallengeNotify, func() any { return new(proto.TreasureMapBonusChallengeNotify) })
	c.regMsg(TreasureMapCurrencyNotify, func() any { return new(proto.TreasureMapCurrencyNotify) })
	c.regMsg(TreasureMapDetectorDataNotify, func() any { return new(proto.TreasureMapDetectorDataNotify) })
	c.regMsg(TreasureMapGuideTaskDoneNotify, func() any { return new(proto.TreasureMapGuideTaskDoneNotify) })
	c.regMsg(TreasureMapHostInfoNotify, func() any { return new(proto.TreasureMapHostInfoNotify) })
	c.regMsg(TreasureMapMpChallengeNotify, func() any { return new(proto.TreasureMapMpChallengeNotify) })
	c.regMsg(TreasureMapPreTaskDoneNotify, func() any { return new(proto.TreasureMapPreTaskDoneNotify) })
	c.regMsg(TreasureMapRegionActiveNotify, func() any { return new(proto.TreasureMapRegionActiveNotify) })
	c.regMsg(TreasureMapRegionInfoNotify, func() any { return new(proto.TreasureMapRegionInfoNotify) })
	c.regMsg(TreasureSeelieCollectOrbsNotify, func() any { return new(proto.TreasureSeelieCollectOrbsNotify) })
	c.regMsg(TrialAvatarFirstPassDungeonNotify, func() any { return new(proto.TrialAvatarFirstPassDungeonNotify) })
	c.regMsg(TrialAvatarInDungeonIndexNotify, func() any { return new(proto.TrialAvatarInDungeonIndexNotify) })
	c.regMsg(TriggerCreateGadgetToEquipPartNotify, func() any { return new(proto.TriggerCreateGadgetToEquipPartNotify) })
	c.regMsg(TriggerRoguelikeCurseNotify, func() any { return new(proto.TriggerRoguelikeCurseNotify) })
	c.regMsg(TriggerRoguelikeRuneReq, func() any { return new(proto.TriggerRoguelikeRuneReq) })
	c.regMsg(TriggerRoguelikeRuneRsp, func() any { return new(proto.TriggerRoguelikeRuneRsp) })
	c.regMsg(TryCustomDungeonReq, func() any { return new(proto.TryCustomDungeonReq) })
	c.regMsg(TryCustomDungeonRsp, func() any { return new(proto.TryCustomDungeonRsp) })
	c.regMsg(TryEnterHomeReq, func() any { return new(proto.TryEnterHomeReq) })
	c.regMsg(TryEnterHomeRsp, func() any { return new(proto.TryEnterHomeRsp) })
	c.regMsg(TryEnterNextRogueDiaryDungeonReq, func() any { return new(proto.TryEnterNextRogueDiaryDungeonReq) })
	c.regMsg(TryEnterNextRogueDiaryDungeonRsp, func() any { return new(proto.TryEnterNextRogueDiaryDungeonRsp) })
	c.regMsg(TryInterruptRogueDiaryDungeonReq, func() any { return new(proto.TryInterruptRogueDiaryDungeonReq) })
	c.regMsg(TryInterruptRogueDiaryDungeonRsp, func() any { return new(proto.TryInterruptRogueDiaryDungeonRsp) })
	c.regMsg(UgcNotify, func() any { return new(proto.UgcNotify) })
	c.regMsg(UnfreezeGroupLimitNotify, func() any { return new(proto.UnfreezeGroupLimitNotify) })
	c.regMsg(UnionCmdNotify, func() any { return new(proto.UnionCmdNotify) })
	c.regMsg(UnlockAvatarTalentReq, func() any { return new(proto.UnlockAvatarTalentReq) })
	c.regMsg(UnlockAvatarTalentRsp, func() any { return new(proto.UnlockAvatarTalentRsp) })
	c.regMsg(UnlockCoopChapterReq, func() any { return new(proto.UnlockCoopChapterReq) })
	c.regMsg(UnlockCoopChapterRsp, func() any { return new(proto.UnlockCoopChapterRsp) })
	c.regMsg(UnlockNameCardNotify, func() any { return new(proto.UnlockNameCardNotify) })
	c.regMsg(UnlockPersonalLineReq, func() any { return new(proto.UnlockPersonalLineReq) })
	c.regMsg(UnlockPersonalLineRsp, func() any { return new(proto.UnlockPersonalLineRsp) })
	c.regMsg(UnlockTransPointReq, func() any { return new(proto.UnlockTransPointReq) })
	c.regMsg(UnlockTransPointRsp, func() any { return new(proto.UnlockTransPointRsp) })
	c.regMsg(UnlockedFurnitureFormulaDataNotify, func() any { return new(proto.UnlockedFurnitureFormulaDataNotify) })
	c.regMsg(UnlockedFurnitureSuiteDataNotify, func() any { return new(proto.UnlockedFurnitureSuiteDataNotify) })
	c.regMsg(UnmarkEntityInMinMapNotify, func() any { return new(proto.UnmarkEntityInMinMapNotify) })
	c.regMsg(UpdateAbilityCreatedMovingPlatformNotify, func() any { return new(proto.UpdateAbilityCreatedMovingPlatformNotify) })
	c.regMsg(UpdatePS4BlockListReq, func() any { return new(proto.UpdatePS4BlockListReq) })
	c.regMsg(UpdatePS4BlockListRsp, func() any { return new(proto.UpdatePS4BlockListRsp) })
	c.regMsg(UpdatePS4FriendListNotify, func() any { return new(proto.UpdatePS4FriendListNotify) })
	c.regMsg(UpdatePS4FriendListReq, func() any { return new(proto.UpdatePS4FriendListReq) })
	c.regMsg(UpdatePS4FriendListRsp, func() any { return new(proto.UpdatePS4FriendListRsp) })
	c.regMsg(UpdatePlayerShowAvatarListReq, func() any { return new(proto.UpdatePlayerShowAvatarListReq) })
	c.regMsg(UpdatePlayerShowAvatarListRsp, func() any { return new(proto.UpdatePlayerShowAvatarListRsp) })
	c.regMsg(UpdatePlayerShowNameCardListReq, func() any { return new(proto.UpdatePlayerShowNameCardListReq) })
	c.regMsg(UpdatePlayerShowNameCardListRsp, func() any { return new(proto.UpdatePlayerShowNameCardListRsp) })
	c.regMsg(UpdateRedPointNotify, func() any { return new(proto.UpdateRedPointNotify) })
	c.regMsg(UpdateReunionWatcherNotify, func() any { return new(proto.UpdateReunionWatcherNotify) })
	c.regMsg(UpdateSalvageBundleMarkReq, func() any { return new(proto.UpdateSalvageBundleMarkReq) })
	c.regMsg(UpdateSalvageBundleMarkRsp, func() any { return new(proto.UpdateSalvageBundleMarkRsp) })
	c.regMsg(UpgradeRoguelikeShikigamiReq, func() any { return new(proto.UpgradeRoguelikeShikigamiReq) })
	c.regMsg(UpgradeRoguelikeShikigamiRsp, func() any { return new(proto.UpgradeRoguelikeShikigamiRsp) })
	c.regMsg(UseItemReq, func() any { return new(proto.UseItemReq) })
	c.regMsg(UseItemRsp, func() any { return new(proto.UseItemRsp) })
	c.regMsg(UseMiracleRingReq, func() any { return new(proto.UseMiracleRingReq) })
	c.regMsg(UseMiracleRingRsp, func() any { return new(proto.UseMiracleRingRsp) })
	c.regMsg(UseWidgetCreateGadgetReq, func() any { return new(proto.UseWidgetCreateGadgetReq) })
	c.regMsg(UseWidgetCreateGadgetRsp, func() any { return new(proto.UseWidgetCreateGadgetRsp) })
	c.regMsg(UseWidgetRetractGadgetReq, func() any { return new(proto.UseWidgetRetractGadgetReq) })
	c.regMsg(UseWidgetRetractGadgetRsp, func() any { return new(proto.UseWidgetRetractGadgetRsp) })
	c.regMsg(VehicleInteractReq, func() any { return new(proto.VehicleInteractReq) })
	c.regMsg(VehicleInteractRsp, func() any { return new(proto.VehicleInteractRsp) })
	c.regMsg(VehicleStaminaNotify, func() any { return new(proto.VehicleStaminaNotify) })
	c.regMsg(ViewCodexReq, func() any { return new(proto.ViewCodexReq) })
	c.regMsg(ViewCodexRsp, func() any { return new(proto.ViewCodexRsp) })
	c.regMsg(ViewLanternProjectionLevelTipsReq, func() any { return new(proto.ViewLanternProjectionLevelTipsReq) })
	c.regMsg(ViewLanternProjectionLevelTipsRsp, func() any { return new(proto.ViewLanternProjectionLevelTipsRsp) })
	c.regMsg(ViewLanternProjectionTipsReq, func() any { return new(proto.ViewLanternProjectionTipsReq) })
	c.regMsg(ViewLanternProjectionTipsRsp, func() any { return new(proto.ViewLanternProjectionTipsRsp) })
	c.regMsg(VintageCampGroupBundleRegisterNotify, func() any { return new(proto.VintageCampGroupBundleRegisterNotify) })
	c.regMsg(VintageCampStageFinishNotify, func() any { return new(proto.VintageCampStageFinishNotify) })
	c.regMsg(VintageDecorateBoothReq, func() any { return new(proto.VintageDecorateBoothReq) })
	c.regMsg(VintageDecorateBoothRsp, func() any { return new(proto.VintageDecorateBoothRsp) })
	c.regMsg(VintageHuntingStartGalleryReq, func() any { return new(proto.VintageHuntingStartGalleryReq) })
	c.regMsg(VintageHuntingStartGalleryRsp, func() any { return new(proto.VintageHuntingStartGalleryRsp) })
	c.regMsg(VintageMarketDeliverItemReq, func() any { return new(proto.VintageMarketDeliverItemReq) })
	c.regMsg(VintageMarketDeliverItemRsp, func() any { return new(proto.VintageMarketDeliverItemRsp) })
	c.regMsg(VintageMarketDividendFinishNotify, func() any { return new(proto.VintageMarketDividendFinishNotify) })
	c.regMsg(VintageMarketFinishStorePlayReq, func() any { return new(proto.VintageMarketFinishStorePlayReq) })
	c.regMsg(VintageMarketFinishStorePlayRsp, func() any { return new(proto.VintageMarketFinishStorePlayRsp) })
	c.regMsg(VintageMarketNpcEventFinishNotify, func() any { return new(proto.VintageMarketNpcEventFinishNotify) })
	c.regMsg(VintageMarketStartStorePlayReq, func() any { return new(proto.VintageMarketStartStorePlayReq) })
	c.regMsg(VintageMarketStartStorePlayRsp, func() any { return new(proto.VintageMarketStartStorePlayRsp) })
	c.regMsg(VintageMarketStoreChooseStrategyReq, func() any { return new(proto.VintageMarketStoreChooseStrategyReq) })
	c.regMsg(VintageMarketStoreChooseStrategyRsp, func() any { return new(proto.VintageMarketStoreChooseStrategyRsp) })
	c.regMsg(VintageMarketStoreUnlockSlotReq, func() any { return new(proto.VintageMarketStoreUnlockSlotReq) })
	c.regMsg(VintageMarketStoreUnlockSlotRsp, func() any { return new(proto.VintageMarketStoreUnlockSlotRsp) })
	c.regMsg(VintageMarketStoreViewStrategyReq, func() any { return new(proto.VintageMarketStoreViewStrategyReq) })
	c.regMsg(VintageMarketStoreViewStrategyRsp, func() any { return new(proto.VintageMarketStoreViewStrategyRsp) })
	c.regMsg(VintagePresentFinishNoify, func() any { return new(proto.VintagePresentFinishNoify) })
	c.regMsg(VintagePresentFinishNotify, func() any { return new(proto.VintagePresentFinishNotify) })
	c.regMsg(WatcherAllDataNotify, func() any { return new(proto.WatcherAllDataNotify) })
	c.regMsg(WatcherChangeNotify, func() any { return new(proto.WatcherChangeNotify) })
	c.regMsg(WatcherEventNotify, func() any { return new(proto.WatcherEventNotify) })
	c.regMsg(WatcherEventStageNotify, func() any { return new(proto.WatcherEventStageNotify) })
	c.regMsg(WatcherEventTypeNotify, func() any { return new(proto.WatcherEventTypeNotify) })
	c.regMsg(WaterSpritePhaseFinishNotify, func() any { return new(proto.WaterSpritePhaseFinishNotify) })
	c.regMsg(WeaponAwakenReq, func() any { return new(proto.WeaponAwakenReq) })
	c.regMsg(WeaponAwakenRsp, func() any { return new(proto.WeaponAwakenRsp) })
	c.regMsg(WeaponPromoteReq, func() any { return new(proto.WeaponPromoteReq) })
	c.regMsg(WeaponPromoteRsp, func() any { return new(proto.WeaponPromoteRsp) })
	c.regMsg(WeaponUpgradeReq, func() any { return new(proto.WeaponUpgradeReq) })
	c.regMsg(WeaponUpgradeRsp, func() any { return new(proto.WeaponUpgradeRsp) })
	c.regMsg(WearEquipReq, func() any { return new(proto.WearEquipReq) })
	c.regMsg(WearEquipRsp, func() any { return new(proto.WearEquipRsp) })
	c.regMsg(WidgetActiveChangeNotify, func() any { return new(proto.WidgetActiveChangeNotify) })
	c.regMsg(WidgetCaptureAnimalReq, func() any { return new(proto.WidgetCaptureAnimalReq) })
	c.regMsg(WidgetCaptureAnimalRsp, func() any { return new(proto.WidgetCaptureAnimalRsp) })
	c.regMsg(WidgetCoolDownNotify, func() any { return new(proto.WidgetCoolDownNotify) })
	c.regMsg(WidgetDoBagReq, func() any { return new(proto.WidgetDoBagReq) })
	c.regMsg(WidgetDoBagRsp, func() any { return new(proto.WidgetDoBagRsp) })
	c.regMsg(WidgetGadgetAllDataNotify, func() any { return new(proto.WidgetGadgetAllDataNotify) })
	c.regMsg(WidgetGadgetDataNotify, func() any { return new(proto.WidgetGadgetDataNotify) })
	c.regMsg(WidgetGadgetDestroyNotify, func() any { return new(proto.WidgetGadgetDestroyNotify) })
	c.regMsg(WidgetQuickHitTreeReq, func() any { return new(proto.WidgetQuickHitTreeReq) })
	c.regMsg(WidgetQuickHitTreeRsp, func() any { return new(proto.WidgetQuickHitTreeRsp) })
	c.regMsg(WidgetReportReq, func() any { return new(proto.WidgetReportReq) })
	c.regMsg(WidgetReportRsp, func() any { return new(proto.WidgetReportRsp) })
	c.regMsg(WidgetSlotChangeNotify, func() any { return new(proto.WidgetSlotChangeNotify) })
	c.regMsg(WidgetUpdateExtraCDReq, func() any { return new(proto.WidgetUpdateExtraCDReq) })
	c.regMsg(WidgetUpdateExtraCDRsp, func() any { return new(proto.WidgetUpdateExtraCDRsp) })
	c.regMsg(WidgetUseAttachAbilityGroupChangeNotify, func() any { return new(proto.WidgetUseAttachAbilityGroupChangeNotify) })
	c.regMsg(WindFieldGalleryChallengeInfoNotify, func() any { return new(proto.WindFieldGalleryChallengeInfoNotify) })
	c.regMsg(WindFieldGalleryInfoNotify, func() any { return new(proto.WindFieldGalleryInfoNotify) })
	c.regMsg(WindFieldRestartDungeonReq, func() any { return new(proto.WindFieldRestartDungeonReq) })
	c.regMsg(WindFieldRestartDungeonRsp, func() any { return new(proto.WindFieldRestartDungeonRsp) })
	c.regMsg(WindSeedClientNotify, func() any { return new(proto.WindSeedClientNotify) })
	c.regMsg(WinterCampAcceptAllGiveItemReq, func() any { return new(proto.WinterCampAcceptAllGiveItemReq) })
	c.regMsg(WinterCampAcceptAllGiveItemRsp, func() any { return new(proto.WinterCampAcceptAllGiveItemRsp) })
	c.regMsg(WinterCampAcceptGiveItemReq, func() any { return new(proto.WinterCampAcceptGiveItemReq) })
	c.regMsg(WinterCampAcceptGiveItemRsp, func() any { return new(proto.WinterCampAcceptGiveItemRsp) })
	c.regMsg(WinterCampEditSnowmanCombinationReq, func() any { return new(proto.WinterCampEditSnowmanCombinationReq) })
	c.regMsg(WinterCampEditSnowmanCombinationRsp, func() any { return new(proto.WinterCampEditSnowmanCombinationRsp) })
	c.regMsg(WinterCampGetCanGiveFriendItemReq, func() any { return new(proto.WinterCampGetCanGiveFriendItemReq) })
	c.regMsg(WinterCampGetCanGiveFriendItemRsp, func() any { return new(proto.WinterCampGetCanGiveFriendItemRsp) })
	c.regMsg(WinterCampGetFriendWishListReq, func() any { return new(proto.WinterCampGetFriendWishListReq) })
	c.regMsg(WinterCampGetFriendWishListRsp, func() any { return new(proto.WinterCampGetFriendWishListRsp) })
	c.regMsg(WinterCampGetRecvItemListReq, func() any { return new(proto.WinterCampGetRecvItemListReq) })
	c.regMsg(WinterCampGetRecvItemListRsp, func() any { return new(proto.WinterCampGetRecvItemListRsp) })
	c.regMsg(WinterCampGiveFriendItemReq, func() any { return new(proto.WinterCampGiveFriendItemReq) })
	c.regMsg(WinterCampGiveFriendItemRsp, func() any { return new(proto.WinterCampGiveFriendItemRsp) })
	c.regMsg(WinterCampRaceScoreNotify, func() any { return new(proto.WinterCampRaceScoreNotify) })
	c.regMsg(WinterCampRecvItemNotify, func() any { return new(proto.WinterCampRecvItemNotify) })
	c.regMsg(WinterCampSetWishListReq, func() any { return new(proto.WinterCampSetWishListReq) })
	c.regMsg(WinterCampSetWishListRsp, func() any { return new(proto.WinterCampSetWishListRsp) })
	c.regMsg(WinterCampStageInfoChangeNotify, func() any { return new(proto.WinterCampStageInfoChangeNotify) })
	c.regMsg(WinterCampTakeBattleRewardReq, func() any { return new(proto.WinterCampTakeBattleRewardReq) })
	c.regMsg(WinterCampTakeBattleRewardRsp, func() any { return new(proto.WinterCampTakeBattleRewardRsp) })
	c.regMsg(WinterCampTakeExploreRewardReq, func() any { return new(proto.WinterCampTakeExploreRewardReq) })
	c.regMsg(WinterCampTakeExploreRewardRsp, func() any { return new(proto.WinterCampTakeExploreRewardRsp) })
	c.regMsg(WinterCampTriathlonRestartReq, func() any { return new(proto.WinterCampTriathlonRestartReq) })
	c.regMsg(WinterCampTriathlonRestartRsp, func() any { return new(proto.WinterCampTriathlonRestartRsp) })
	c.regMsg(WinterCampTriathlonSettleNotify, func() any { return new(proto.WinterCampTriathlonSettleNotify) })
	c.regMsg(WorktopOptionNotify, func() any { return new(proto.WorktopOptionNotify) })
	c.regMsg(WorldAllRoutineTypeNotify, func() any { return new(proto.WorldAllRoutineTypeNotify) })
	c.regMsg(WorldChestOpenNotify, func() any { return new(proto.WorldChestOpenNotify) })
	c.regMsg(WorldDataNotify, func() any { return new(proto.WorldDataNotify) })
	c.regMsg(WorldOwnerBlossomBriefInfoNotify, func() any { return new(proto.WorldOwnerBlossomBriefInfoNotify) })
	c.regMsg(WorldOwnerBlossomScheduleInfoNotify, func() any { return new(proto.WorldOwnerBlossomScheduleInfoNotify) })
	c.regMsg(WorldOwnerDailyTaskNotify, func() any { return new(proto.WorldOwnerDailyTaskNotify) })
	c.regMsg(WorldPlayerDieNotify, func() any { return new(proto.WorldPlayerDieNotify) })
	c.regMsg(WorldPlayerInfoNotify, func() any { return new(proto.WorldPlayerInfoNotify) })
	c.regMsg(WorldPlayerLocationNotify, func() any { return new(proto.WorldPlayerLocationNotify) })
	c.regMsg(WorldPlayerRTTNotify, func() any { return new(proto.WorldPlayerRTTNotify) })
	c.regMsg(WorldPlayerReviveReq, func() any { return new(proto.WorldPlayerReviveReq) })
	c.regMsg(WorldPlayerReviveRsp, func() any { return new(proto.WorldPlayerReviveRsp) })
	c.regMsg(WorldRoutineChangeNotify, func() any { return new(proto.WorldRoutineChangeNotify) })
	c.regMsg(WorldRoutineTypeCloseNotify, func() any { return new(proto.WorldRoutineTypeCloseNotify) })
	c.regMsg(WorldRoutineTypeRefreshNotify, func() any { return new(proto.WorldRoutineTypeRefreshNotify) })
}

const (
	GCGDSBanCardNotify                = 65001
	GCGStartChallengeByCheckRewardReq = 65002
	GCGStartChallengeByCheckRewardRsp = 65003
)
