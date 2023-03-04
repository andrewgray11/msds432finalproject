DROP DATABASE IF EXISTS chicagobi;
CREATE DATABASE chicagobi;

DROP TABLE IF EXISTS buildingPermit;
DROP TABLE IF EXISTS ccvi;
DROP TABLE IF EXISTS censusData;
DROP TABLE IF EXISTS communityBound;
DROP TABLE IF EXISTS covidCasesZip;
/*DROP TABLE IF EXISTS covidDaily;*/
DROP TABLE IF EXISTS pubHealthStats;
DROP TABLE IF EXISTS taxiTrips;
DROP TABLE IF EXISTS transNetProviderTrips;

\c chicagobi;

CREATE TABLE buildingPermit (

    	ID                  INT,    
	PermitNumber        VARCHAR(100),
	PermitType          VARCHAR(100),
	ReviewType          VARCHAR(100),
	TotalFee	    FLOAT,
	AppStartDate        VARCHAR(100),
	IssueDate	    VARCHAR(100),
	CommunityArea       INT,
	Latitude	    FLOAT,
	Longitude    	    FLOAT,
);

CREATE TABLE ccvi (

	GeographyType        VARCHAR(100),
	CommunityAreaOrZip   INT,
	CommunityAreaName    VARCHAR(100), 
	CcviCategory         VARCHAR(100), 
	CcviScore            FLOAT, 
	SocioeconomicStatus  INT, 
	HouseholdComposition INT, 
	NoPrimaryCare        INT, 
	CumMobilityRatio     INT, 
	FrontlineWorkers     INT, 
	Age65OrGreater       INT, 
	ComorbidConditions   INT, 
	CovidIncidenceRate   INT, 
	CovidHospitalRate    INT, 
	CrudeMortalityRate   INT
);

CREATE TABLE censusData (

	CommunityAreaNumber                   INT,
	CommunityAreaName                     VARCHAR(100),
	PercentOfHousingCrowded               FLOAT, 
	PercentHouseholdsBelowPoverty         FLOAT, 
	PercentAged16Unemployed               FLOAT, 
	PercentAged25WithoutHighSchoolDiploma FLOAT,
	PercentAgedUnder18OrOver64            FLOAT,
	PerCapitaIncome                       FLOAT, 
	HardshipIndex                         INT
);

CREATE TABLE communityBound (

	AreaNumber                            INT,
	AreaName                              VARCHAR(100),
	Geometry                              VARCHAR(100) 
	
);

CREATE TABLE covidCasesZip (

	Zip          VARCHAR(100),
	WeekNumber   INT,
	WeekStart    VARCHAR(100),
	WeekEnd      VARCHAR(100),
	CasesWeekly  INT,
	CasesMonthly INT,
	CasesCum     INT,
	CaseRateWkly FLOAT,
	CaseRateMnth FLOAT,
	CaseRateCum  FLOAT,
	TestsWeekly  INT,
	TestsCum     INT,
	TestRateWkly FLOAT,
	TestRateCum  FLOAT,
	PosRateWkly  FLOAT,
	PosRateCum   FLOAT,
	DeathsWkly   INT,
	DeathsCum    INT,
	DeathRtWkly  FLOAT,
	DeathRtCum   FLOAT,
	Population   INT,
	RowId        VARCHAR(100)
);

/*CREATE TABLE covidDaily (

	Date                VARCHAR(20),
	TotalCases          VARCHAR(20),
	NewCases            VARCHAR(20),
	TotalDeaths         VARCHAR(20),
	NewDeaths           VARCHAR(20),
	TotalHospitalized   VARCHAR(20),
	NewHospitalized     VARCHAR(20),
	AverageDailyHospitalized VARCHAR(20)

);*/

CREATE TABLE pubHealthStats (

	CommunityArea                         VARCHAR(100),
	CommunityAreaName                     VARCHAR(100),
	BirthRate                             VARCHAR(100),
	GeneralFertilityRate                  VARCHAR(100),
	LowBirthWeight                        VARCHAR(100),
	PrenatalCareBeginningInFirstTrimester FLOAT,
	PretermBirths                         VARCHAR(100),
	TeenBirthRate                         FLOAT,
	AssaultHomicide                       VARCHAR(100),
	BreastCancerInFemales                 FLOAT,
	CancerAllSites                        VARCHAR(100),
	ColorectalCancer                      FLOAT,
	DiabetesRelated                       VARCHAR(100),
	FirearmRelated                        FLOAT,
	InfantMortalityRate                   FLOAT,
	LungCancer                            FLOAT,
	ProstateCancerInMales                 FLOAT,
	StrokeCerebrovascularDisease          FLOAT,
	ChildhoodBloodLeadLevelScreening      FLOAT,
	ChildhoodLeadPoisoning                FLOAT,
	GonorrheaInFemales                    FLOAT,
	GonorrheaInMales                      VARCHAR(100),
	Tuberculosis                          FLOAT, 
	BelowPovertyLevel                     FLOAT,
	CrowdedHousing                        FLOAT,
	Dependency                            FLOAT,
	NoHighSchoolDiploma                   FLOAT,
	PerCapitaIncome                       FLOAT,
	Unemployment                          FLOAT
);

CREATE TABLE taxiTrips (

    TripID                   VARCHAR(100),
    TaxiID                   VARCHAR(100),
    TripStartTimestamp       TIMESTAMP,
    TripEndTimestamp         TIMESTAMP,
    TripSeconds              INT,
    TripMiles                FLOAT,
    PickupCensusTract        VARCHAR(100),
    DropoffCensusTract       VARCHAR(100),
    PickupCommunityArea      INT,
    DropoffCommunityArea     INT,
    Fare                     FLOAT,
    Tips                     FLOAT,
    Tolls                    FLOAT,
    Extras                   FLOAT,
    TripTotal                FLOAT,
    PaymentType              VARCHAR(100),
    Company                  VARCHAR(100),
    PickupCentroidLatitude   FLOAT,
    PickupCentroidLongitude  FLOAT,
    PickupCentroidLocation   VARCHAR(100),
    DropoffCentroidLatitude  FLOAT,
    DropoffCentroidLongitude FLOAT,
    DropoffCentroidLocation  VARCHAR(100)

);

CREATE TABLE transNetProviderTrips (
    TripID                   VARCHAR(100),
    TripStartTimestamp       VARCHAR(100),
    TripEndTimestamp         VARCHAR(100),
    TripSeconds              INT,
    TripMiles                FLOAT,
    PickupCensusTract        VARCHAR(100),
    DropoffCensusTract       VARCHAR(100),
    PickupCommunityArea      INT,
    DropoffCommunityArea     INT,
    Fare                     FLOAT,
    Tip                      FLOAT,
    AdditionalCharges        FLOAT,
    TripTotal                FLOAT,
    SharedTripAuthorized     BOOLEAN,
    TripsPooled              INT,
    PickupCentroidLatitude   FLOAT,
    PickupCentroidLongitude  FLOAT,
    PickupCentroidLocation   VARCHAR(100),
    DropoffCentroidLatitude  FLOAT,
    DropoffCentroidLongitude FLOAT,
    DropoffCentroidLocation  VARCHAR(100)
);
