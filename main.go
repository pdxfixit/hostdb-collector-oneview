package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/pdxfixit/hostdb"
)

func main() {

	// load config
	loadConfig()

	for _, oneviewHost := range config.OneView.Hosts {

		var clientOV *ov.OVClient

		// get a session token
		ovc := clientOV.NewOVClient(
			config.OneView.User,
			config.OneView.Pass,
			config.OneView.Domain,
			oneviewHost,
			false,
			1200,
			"*")

		log.Println(fmt.Sprintf(
			"Accessing %s : %s @ %s...",
			config.OneView.User,
			"*****"+config.OneView.Pass[len(config.OneView.Pass)-3:],
			oneviewHost,
		))

		//
		// enclosures
		//
		enclosures, err := ovc.GetEnclosures("", "", "", "name:desc", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var enclosureRecords []hostdb.Record
		for _, e := range enclosures.Members {
			data, err := json.Marshal(e)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			enclosureRecords = append(enclosureRecords, hostdb.Record{
				Type: "oneview-enclosure",
				Data: data,
			})
		}

		if err := postToHostdb(enclosureRecords, "oneview-enclosure", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// enclosure groups
		//
		enclosureGroups, err := ovc.GetEnclosureGroups("", "", "", "name:desc", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var enclosureGroupRecords []hostdb.Record
		for _, eg := range enclosureGroups.Members {
			data, err := json.Marshal(eg)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			enclosureGroupRecords = append(enclosureGroupRecords, hostdb.Record{
				Type: "oneview-enclosure_group",
				Data: data,
			})
		}

		if err := postToHostdb(enclosureGroupRecords, "oneview-enclosure_group", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// ethernet networks
		//
		ethernetNetworks, err := ovc.GetEthernetNetworks("", "", "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var ethernetNetworkRecords []hostdb.Record
		for _, eg := range ethernetNetworks.Members {
			data, err := json.Marshal(eg)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			ethernetNetworkRecords = append(ethernetNetworkRecords, hostdb.Record{
				Type: "oneview-ethernet_network",
				Data: data,
			})
		}

		if err := postToHostdb(ethernetNetworkRecords, "oneview-ethernet_network", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// fc networks
		//
		fcNetworks, err := ovc.GetFCNetworks("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var fcNetworkRecords []hostdb.Record
		for _, fcn := range fcNetworks.Members {
			data, err := json.Marshal(fcn)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			fcNetworkRecords = append(fcNetworkRecords, hostdb.Record{
				Type: "oneview-fc_network",
				Data: data,
			})
		}

		if err := postToHostdb(fcNetworkRecords, "oneview-fc_network", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// fcoe networks
		//
		fcoeNetworks, err := ovc.GetFCoENetworks("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var fcoeNetworkRecords []hostdb.Record
		for _, fcoeNetwork := range fcoeNetworks.Members {
			data, err := json.Marshal(fcoeNetwork)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			fcoeNetworkRecords = append(fcoeNetworkRecords, hostdb.Record{
				Type: "oneview-fcoe_network",
				Data: data,
			})
		}

		if err := postToHostdb(fcNetworkRecords, "oneview-fcoe_network", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// interconnects
		//
		interconnects, err := ovc.GetInterconnects("", "", "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var interconnectRecords []hostdb.Record
		for _, interconnect := range interconnects.Members {
			// hostname
			hostname := ""
			if interconnect.HostName != "" && interconnect.HostName != "none" {
				hostname = interconnect.HostName
			}

			// ip
			ip := interconnect.InterconnectIP

			// data
			data, err := json.Marshal(interconnect)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			interconnectRecords = append(interconnectRecords, hostdb.Record{
				Hostname: hostname,
				IP:       ip,
				Type:     "oneview-interconnect",
				Data:     data,
			})
		}

		if err := postToHostdb(interconnectRecords, "oneview-interconnect", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// interconnect type
		//
		interconnectTypes, err := ovc.GetInterconnectTypes("", "", "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var interconnectTypeRecords []hostdb.Record
		for _, interconnectType := range interconnectTypes.Members {
			data, err := json.Marshal(interconnectType)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			interconnectTypeRecords = append(interconnectTypeRecords, hostdb.Record{
				Type: "oneview-interconnect_type",
				Data: data,
			})
		}

		if err := postToHostdb(interconnectTypeRecords, "oneview-interconnect_type", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// logical enclosures
		//
		logicalEnclosures, err := ovc.GetLogicalEnclosures("", "", "", nil, "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var logicalEnclosureRecords []hostdb.Record
		for _, logicalEnclosure := range logicalEnclosures.Members {
			data, err := json.Marshal(logicalEnclosure)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			logicalEnclosureRecords = append(logicalEnclosureRecords, hostdb.Record{
				Type: "oneview-logical_enclosure",
				Data: data,
			})
		}

		if err := postToHostdb(logicalEnclosureRecords, "oneview-logical_enclosure", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// logical interconnect groups
		//
		logicalInterconnectGroups, err := ovc.GetLogicalInterconnectGroups(0, "", "", "name:desc", 0)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var logicalInterconnectGroupRecords []hostdb.Record
		for _, logicalInterconnectGroup := range logicalInterconnectGroups.Members {
			data, err := json.Marshal(logicalInterconnectGroup)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			logicalInterconnectGroupRecords = append(logicalInterconnectGroupRecords, hostdb.Record{
				Type: "oneview-logical_interconnect_group",
				Data: data,
			})
		}

		if err := postToHostdb(logicalInterconnectGroupRecords, "oneview-logical_interconnect_group", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// logical interconnects
		//
		logicalInterconnects, err := ovc.GetLogicalInterconnects("", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var logicalInterconnectRecords []hostdb.Record
		for _, logicalInterconnect := range logicalInterconnects.Members {
			data, err := json.Marshal(logicalInterconnect)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			logicalInterconnectRecords = append(logicalInterconnectRecords, hostdb.Record{
				Type: "oneview-logical_interconnect",
				Data: data,
			})
		}

		if err := postToHostdb(logicalInterconnectRecords, "oneview-logical_interconnect", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// network sets
		//
		networkSets, err := ovc.GetNetworkSets("", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var networkSetRecords []hostdb.Record
		for _, networkSet := range networkSets.Members {
			data, err := json.Marshal(networkSet)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			networkSetRecords = append(networkSetRecords, hostdb.Record{
				Type: "oneview-network_set",
				Data: data,
			})
		}

		if err := postToHostdb(networkSetRecords, "oneview-network_set", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// scopes
		//
		scopes, err := ovc.GetScopes("", "", "", "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var scopeRecords []hostdb.Record
		for _, scope := range scopes.Members {
			data, err := json.Marshal(scope)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			scopeRecords = append(scopeRecords, hostdb.Record{
				Type: "oneview-scope",
				Data: data,
			})
		}

		if err := postToHostdb(scopeRecords, "oneview-scope", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// server hardware
		//
		serverHardware, err := ovc.GetServerHardwareList([]string{}, "", "", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var serverHardwareRecords []hostdb.Record
		for _, sh := range serverHardware.Members {

			// hostname
			hostname := sh.MpHostInfo.MpHostName

			// ip
			ip := ""
			for _, a := range sh.MpHostInfo.MpIPAddresses {
				if a.Type != "Static" {
					continue
				}

				ip = a.Address
			}

			// data
			data, err := json.Marshal(sh)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			serverHardwareRecords = append(serverHardwareRecords, hostdb.Record{
				Hostname: hostname,
				IP:       ip,
				Type:     "oneview-server_hardware",
				Data:     data,
			})
		}

		if err := postToHostdb(serverHardwareRecords, "oneview-server_hardware", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// server hardware types
		//
		serverHardwareType, err := ovc.GetServerHardwareTypes(0, 0, "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var serverHardwareTypeRecords []hostdb.Record
		for _, shType := range serverHardwareType.Members {
			data, err := json.Marshal(shType)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			serverHardwareTypeRecords = append(serverHardwareTypeRecords, hostdb.Record{
				Type: "oneview-server_hardware_type",
				Data: data,
			})
		}

		if err := postToHostdb(serverHardwareTypeRecords, "oneview-server_hardware_type", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// server profiles
		//
		serverProfiles, err := ovc.GetProfiles("", "", "", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var serverProfileRecords []hostdb.Record
		for _, sp := range serverProfiles.Members {
			data, err := json.Marshal(sp)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			serverProfileRecords = append(serverProfileRecords, hostdb.Record{
				Type: "oneview-server_profile",
				Data: data,
			})
		}

		if err := postToHostdb(serverProfileRecords, "oneview-server_profile", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// server profile templates
		//
		serverProfileTemplates, err := ovc.GetProfileTemplates("", "", "", "name:desc", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var serverProfileTemplateRecords []hostdb.Record
		for _, spt := range serverProfileTemplates.Members {
			data, err := json.Marshal(spt)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			serverProfileTemplateRecords = append(serverProfileTemplateRecords, hostdb.Record{
				Type: "oneview-server_profile_template",
				Data: data,
			})
		}

		if err := postToHostdb(serverProfileTemplateRecords, "oneview-server_profile_template", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// storage pool
		//
		storagePools, err := ovc.GetStoragePools("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var storagePoolRecords []hostdb.Record
		for _, storagePool := range storagePools.Members {
			data, err := json.Marshal(storagePool)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			storagePoolRecords = append(storagePoolRecords, hostdb.Record{
				Type: "oneview-storage_pool",
				Data: data,
			})
		}

		if err := postToHostdb(storagePoolRecords, "oneview-storage_pool", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// storage system
		//
		storageSystems, err := ovc.GetStorageSystems("", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var storageSystemRecords []hostdb.Record
		for _, storageSystem := range storageSystems.Members {
			data, err := json.Marshal(storageSystem)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			storageSystemRecords = append(storageSystemRecords, hostdb.Record{
				Type: "oneview-storage_system",
				Data: data,
			})
		}

		if err := postToHostdb(storageSystemRecords, "oneview-storage_system", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// storage volume
		//
		storageVolumes, err := ovc.GetStorageVolumes("", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var storageVolumeRecords []hostdb.Record
		for _, storageVolume := range storageVolumes.Members {
			data, err := json.Marshal(storageVolume)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			storageVolumeRecords = append(storageVolumeRecords, hostdb.Record{
				Type: "oneview-storage_volume",
				Data: data,
			})
		}

		if err := postToHostdb(storageVolumeRecords, "oneview-storage_volume", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// storage volume attachment
		//
		storageVolumeAttachments, err := ovc.GetStorageAttachments("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var storageVolumeAttachmentRecords []hostdb.Record
		for _, storageVolumeAttachment := range storageVolumeAttachments.Members {
			data, err := json.Marshal(storageVolumeAttachment)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			storageVolumeAttachmentRecords = append(storageVolumeAttachmentRecords, hostdb.Record{
				Type: "oneview-storage_volume_attachment",
				Data: data,
			})
		}

		if err := postToHostdb(storageVolumeAttachmentRecords, "oneview-storage_volume_attachment", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// storage volume template
		//
		storageVolumeTemplates, err := ovc.GetStorageVolumeTemplates("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var storageVolumeTemplateRecords []hostdb.Record
		for _, storageVolumeTemplate := range storageVolumeTemplates.Members {
			data, err := json.Marshal(storageVolumeTemplate)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			storageVolumeTemplateRecords = append(storageVolumeTemplateRecords, hostdb.Record{
				Type: "oneview-storage_volume_template",
				Data: data,
			})
		}

		if err := postToHostdb(storageVolumeTemplateRecords, "oneview-storage_volume_template", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// tasks
		//
		tasks, err := ovc.GetTasks("", "name:desc", "", "")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var taskRecords []hostdb.Record
		for _, task := range tasks.Members {
			data, err := json.Marshal(task)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			taskRecords = append(taskRecords, hostdb.Record{
				Type: "oneview-task",
				Data: data,
			})
		}

		if err := postToHostdb(taskRecords, "oneview-task", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// uplink set
		//
		uplinkSets, err := ovc.GetUplinkSets("", "", "", "name:desc")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var uplinkSetRecords []hostdb.Record
		for _, uplinkSet := range uplinkSets.Members {
			data, err := json.Marshal(uplinkSet)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			uplinkSetRecords = append(uplinkSetRecords, hostdb.Record{
				Type: "oneview-uplink_set",
				Data: data,
			})
		}

		if err := postToHostdb(uplinkSetRecords, "oneview-uplink_set", oneviewHost); err != nil {
			log.Println(err.Error())
		}

		//
		// DONE
		//
		log.Printf("%s is all done!\n", oneviewHost)

	}

}

func postToHostdb(records []hostdb.Record, recordType string, oneviewHost string) (err error) {

	recordSet := hostdb.RecordSet{
		Type:      recordType,
		Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05"),
		Context: map[string]interface{}{
			"oneview_url": oneviewHost,
		},
		Committer: "hostdb-collector-oneview",
		Records:   records,
	}

	if config.Collector.SampleData {
		if err := recordSet.Save(fmt.Sprintf("%s/%s_%s.json", config.Collector.SampleDataPath, oneviewHost, recordType)); err != nil {
			return err
		}
	} else {
		if err := recordSet.Send(fmt.Sprintf("host=%s&type=%s", oneviewHost, recordType)); err != nil {
			return err
		}
	}

	return nil
}
